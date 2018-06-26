package globaldb

import (
	"os"

	imageIntegrationStore "bitbucket.org/stack-rox/apollo/central/imageintegration/store"
	notifierStore "bitbucket.org/stack-rox/apollo/central/notifier/store"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"github.com/boltdb/bolt"
)

// Export returns a compacted backup of the database.
func Export(db *bolt.DB, exportedFilepath, compactedFilepath string) (*bolt.DB, error) {
	defer os.Remove(exportedFilepath)

	// This will block all other transactions until this has completed. We could use View for a hot backup
	err := db.Update(func(tx *bolt.Tx) error {
		return tx.CopyFile(exportedFilepath, 0600)
	})
	if err != nil {
		return nil, err
	}

	exportDB, err := bolthelper.New(exportedFilepath)
	if err != nil {
		return nil, err
	}
	defer exportDB.Close()

	clearNotifierConfigs(exportDB)
	if err != nil {
		return nil, err
	}

	clearImageIntegrationConfigs(exportDB)
	if err != nil {
		return nil, err
	}

	if err := exportDB.Sync(); err != nil {
		return nil, err
	}

	// Create completely clean DB and compact to it, wiping the secrets from cached memory
	newDB, err := bolt.Open(compactedFilepath, 0600, nil)
	if err != nil {
		return nil, err
	}
	if err := bolthelper.Compact(newDB, exportDB); err != nil {
		return nil, err
	}
	// Close the databases
	exportDB.Close()
	return newDB, nil
}

func clearNotifierConfigs(db *bolt.DB) error {
	store := notifierStore.New(db)

	notifiers, err := store.GetNotifiers(&v1.GetNotifiersRequest{})
	if err != nil {
		return err
	}
	for _, n := range notifiers {
		n.Config = nil
		if err := store.UpdateNotifier(n); err != nil {
			return err
		}
	}
	return nil
}

func clearImageIntegrationConfigs(db *bolt.DB) error {
	store := imageIntegrationStore.New(db)

	integrations, err := store.GetImageIntegrations(&v1.GetImageIntegrationsRequest{})
	if err != nil {
		return err
	}
	for _, d := range integrations {
		d.Config = nil
		if err := store.UpdateImageIntegration(d); err != nil {
			return err
		}
	}
	return nil
}
