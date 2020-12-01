.PHONY: kind-install cluster

TOOL_PATH    := $(PWD)/tmp
CLUSTER_PATH := $(PWD)/deployment/cluster.yaml
DEMO_PATH    := $(PWD)/demo
DEMO_NAME    ?= azure-tanks
TMP_PATH     := $(PWD)/.build

kustomize-install:
	curl -s "https://raw.githubusercontent.com/ \
	kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
	mv ./kustomize $(TOOL_PATH)/

kind-install:
	mkdir -p $(TOOL_PATH)
	curl -Lo $(TOOL_PATH)/kind "https://kind.sigs.k8s.io/dl/v0.9.0/kind-$(uname)-amd64"
	chmod +x $(TOOL_PATH)/kind

demo-build:
	mkdir -p $(TMP_PATH)
	rm -rf $(TMP_PATH)/*
	$(TOOL_PATH)/kustomize build $(DEMO_PATH)/$(DEMO_NAME) > $(TMP_PATH)/deployment.yaml

apply:
	kubectl apply -f $(TMP_PATH)/deployment.yaml --record

delete:
	kubectl delete daemonset northbound-daemon

cluster:
	$(TOOL_PATH)/kind create cluster --name myk8s --config $(CLUSTER_PATH)

clean:
	$(TOOL_PATH)/kind delete cluster --name myk8s

debug:
	docker build \
		-f ./debug.Dockerfile \
		-t moxaisd/logicsim-dev \
		.
	docker create -it --rm \
		--name logicsim \
		-w /data \
		-v ${PWD}:/data \
		moxaisd/logicsim-dev \
		bash
	docker start logicsim
	docker attach logicsim