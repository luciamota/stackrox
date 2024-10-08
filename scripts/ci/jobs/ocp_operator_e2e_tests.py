#!/usr/bin/env -S python3 -u

"""
Run operator e2e tests in an OCP cluster.
"""
import os
from runners import ClusterTestRunner
from clusters import AutomationFlavorsCluster
from ci_tests import OperatorE2eTest
from pre_tests import PreSystemTests
from post_tests import PostClusterTest, FinalPost

os.environ["ORCHESTRATOR_FLAVOR"] = "openshift"

ClusterTestRunner(
    cluster=AutomationFlavorsCluster(),
    pre_test=PreSystemTests(),
    test=OperatorE2eTest(),
    post_test=PostClusterTest(collect_central_artifacts=False),
    final_post=FinalPost(),
).run()
