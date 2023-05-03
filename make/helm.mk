# Copyright 2023 D2iQ, Inc. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

.PHONY: lint-chart
lint-chart: ## Lints helm chart
lint-chart: install-tool.helm-ct install-tool.yamllint
	ct lint --config charts/ct-config.yaml

.PHONY: lint-and-install-chart
lint-and-install-chart: ## Lints and installs helm chart
lint-and-install-chart: install-tool.helm-ct install-tool.yamllint
	ct lint-and-install --config charts/ct-config.yaml
	ct lint-and-install --config charts/ct-config.yaml --upgrade