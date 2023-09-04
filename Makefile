RUN=./scripts/run.sh
MAKE_CONTAINER=$(RUN) make --no-print-directory -e -f Makefile.core.mk
FRONTEND_MAKE_CONTAINER=$(RUN) make --no-print-directory -C frontend -e -f Makefile

%:
	@$(MAKE_CONTAINER) $@

default:
	@$(MAKE_CONTAINER)

shell:
	@$(RUN) /bin/bash

serve:
	@docker-compose up

clean:
	@docker-compose down -v
	@$(MAKE_CONTAINER) clean

.PHONY: default shell clean serve