// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type TestShortCircuitsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestTestShortCircuitsStore(t *testing.T) {
	suite.Run(t, new(TestShortCircuitsStoreSuite))
}

func (s *TestShortCircuitsStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *TestShortCircuitsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE test_short_circuits CASCADE")
	s.T().Log("test_short_circuits", tag)
	s.NoError(err)
}

func (s *TestShortCircuitsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *TestShortCircuitsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	testShortCircuit := &storage.TestShortCircuit{}
	s.NoError(testutils.FullInit(testShortCircuit, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestShortCircuit, exists, err := store.Get(ctx, testShortCircuit.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestShortCircuit)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, testShortCircuit))
	foundTestShortCircuit, exists, err = store.Get(ctx, testShortCircuit.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(testShortCircuit, foundTestShortCircuit)

	testShortCircuitCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, testShortCircuitCount)
	testShortCircuitCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(testShortCircuitCount)

	testShortCircuitExists, err := store.Exists(ctx, testShortCircuit.GetId())
	s.NoError(err)
	s.True(testShortCircuitExists)
	s.NoError(store.Upsert(ctx, testShortCircuit))
	s.ErrorIs(store.Upsert(withNoAccessCtx, testShortCircuit), sac.ErrResourceAccessDenied)

	foundTestShortCircuit, exists, err = store.Get(ctx, testShortCircuit.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(testShortCircuit, foundTestShortCircuit)

	s.NoError(store.Delete(ctx, testShortCircuit.GetId()))
	foundTestShortCircuit, exists, err = store.Get(ctx, testShortCircuit.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestShortCircuit)
	s.NoError(store.Delete(withNoAccessCtx, testShortCircuit.GetId()))

	var testShortCircuits []*storage.TestShortCircuit
	for i := 0; i < 200; i++ {
		testShortCircuit := &storage.TestShortCircuit{}
		s.NoError(testutils.FullInit(testShortCircuit, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		testShortCircuits = append(testShortCircuits, testShortCircuit)
	}

	s.NoError(store.UpsertMany(ctx, testShortCircuits))

	testShortCircuitCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, testShortCircuitCount)
}
