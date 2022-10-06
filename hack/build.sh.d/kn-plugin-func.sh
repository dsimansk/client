#!/usr/bin/env bash

readonly socat_image="${KN_PLUGIN_FUNC_SOCAT_IMAGE:-quay.io/boson/alpine-socat:1.7.4.3-r1-non-root}"
export EXTERNAL_LD_FLAGS="${EXTERNAL_LD_FLAGS:-} \
-X  knative.dev/kn-plugin-func/k8s.socatImage=${socat_image}"

