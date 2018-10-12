bin:
	mkdir -p bin

bin/oc: bin
	./download_oc.sh

.PHONY: vdemeester/oc-builder
vdemeester/oc-builder: bin/oc
	tar cf - bin/oc Dockerfile | docker build -t vdemeester/oc-builder -
