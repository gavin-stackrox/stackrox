include $(CURDIR)/../make/env.mk

ROX_PROJECT=tests
TESTFLAGS=-race -p 1 -timeout 30m

TOPLEVEL=$(CURDIR)/..

.PHONY: all
all: test
	@echo "+ $@"
	@$(MAKE) report JUNIT_OUT=all-tests-results

.PHONY: destructive-tests
destructive-tests:
	@echo "+ $@"
	@GOTAGS=$(GOTAGS),test,destructive ../scripts/go-test.sh -cover -v -run TestClusterDeletion 2>&1 | tee test.log
	@$(MAKE) report JUNIT_OUT=destructive-tests-results

.PHONY: external-backup-tests
external-backup-tests:
	@echo "+ $@"
	@GOTAGS=$(GOTAGS),test,externalbackups ../scripts/go-test.sh -cover -v -run TestGCSExternalBackup 2>&1 | tee test.log
	@$(MAKE) report JUNIT_OUT=external-backup-tests-results

.PHONY: compliance-v2-tests
compliance-v2-tests:
	@GOTAGS=$(GOTAGS),test,compliance ../scripts/go-test.sh -cover -v -run TestComplianceV2 2>&1 | tee test.log
	@$(MAKE) report JUNIT_OUT=compliance-v2-tests-results

include ../make/stackrox.mk
