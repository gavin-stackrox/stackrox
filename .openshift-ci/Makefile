PY_SRCS := *.py **/*.py ../scripts/ci/jobs/*.py
VENV := .dev-venv

ifdef CI
	PYCODESTYLE=pycodestyle
	AUTOPEP8=autopep8
	PYLINT=pylint
	PYTHON3=python3
else
	PYCODESTYLE=$(VENV)/bin/pycodestyle
	AUTOPEP8=$(VENV)/bin/autopep8
	PYLINT=$(VENV)/bin/pylint
	PYTHON3=$(VENV)/bin/python3
endif

.PHONY: all
all: style lint

$(VENV)/bin/activate: dev-requirements.txt
ifdef CI
	$(PYCODESTYLE) --version
	$(PYLINT) --version
else
	python3 -m venv $(VENV)
	$(VENV)/bin/pip3 --quiet --require-virtualenv --no-input --disable-pip-version-check install -r dev-requirements.txt
endif

.PHONY: style
style: $(VENV)/bin/activate
	@# (--max-line-length must be kept in sync with max-line-length in .pylintrc)
	$(PYCODESTYLE) --max-line-length=120 --exclude=$(VENV) $(PY_SRCS)

.PHONY: fix-style
fix-style: $(VENV)/bin/activate
	$(AUTOPEP8) --in-place --exclude=$(VENV) $(PY_SRCS)

.PHONY: lint
lint: $(VENV)/bin/activate
	$(PYLINT) --rcfile .pylintrc --ignore=$(VENV) $(PY_SRCS)

.PHONY: test
test: $(VENV)/bin/activate
	$(PYTHON3) -m unittest discover -v --buffer --start-directory tests
