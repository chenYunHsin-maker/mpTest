#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

cd $GOPATH/src/sdn.io/sdwan/cmd/code-generator

# protobuf
cd go-to-protobuf
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/wifi/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/alert/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1"  --packages="sdn.io/sdwan/pkg/apiserver/apis/batch/v1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1"
go run main.go --apimachinery-packages="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1" --packages="sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1"
go run main.go --packages="sdn.io/sdwan/pkg/apiserver/apis/authentication/v1,sdn.io/sdwan/pkg/apiserver/apis/audit/v1beta1,sdn.io/sdwan/pkg/apiserver/apis/example/v1"

# deepcopy
cd ../deepcopy-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1,sdn.io/sdwan/pkg/apimachinery/apis/meta/v1alpha1,sdn.io/sdwan/pkg/apimachinery/apis/meta/internalversion,sdn.io/sdwan/pkg/apimachinery/runtime,sdn.io/sdwan/pkg/apimachinery/runtime/testing,sdn.io/sdwan/pkg/apimachinery/labels"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apiserver/apis/authentication/v1,sdn.io/sdwan/pkg/apiserver/apis/audit,sdn.io/sdwan/pkg/apiserver/apis/audit/v1beta1,sdn.io/sdwan/pkg/apiserver/apis/example,sdn.io/sdwan/pkg/apiserver/apis/example/v1,sdn.io/sdwan/pkg/apiserver/apis/batch/v1,sdn.io/sdwan/pkg/apiserver/endpoints/testing,sdn.io/sdwan/pkg/apiserver/endpoints/openapi/testing,sdn.io/sdwan/pkg/apiserver/storage/testing"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ifconfig,sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/zoneset,sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/vpn,sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/staticroute,sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/wifi,sdn.io/sdwan/pkg/apis/cubs/wifi/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/webauth,sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alert,sdn.io/sdwan/pkg/apis/cubs/alert/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/firewall,sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/businesspolicy,sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/organization,sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corporation,sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/profile,sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/site,sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/services,sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/clientvpn,sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corptocorp,sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/guiuser,sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/check,sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/serviceclass,sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/troubleshoot,sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/devicedump,sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd,sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/orchpath,sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alg,sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/multicastroute,sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ztp,sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/blackfw,sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1"

# conversion
cd ../conversion-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apiserver/apis/audit/v1beta1,sdn.io/sdwan/pkg/apiserver/apis/example/v1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/wifi/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alert/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1"

# defaulter
cd ../defaulter-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apimachinery/apis/meta/v1,sdn.io/sdwan/pkg/apimachinery/apis/meta/v1alpha1,sdn.io/sdwan/pkg/apimachinery/apis/meta/internalversion,sdn.io/sdwan/pkg/apimachinery/runtime"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apiserver/apis/authentication/v1,sdn.io/sdwan/pkg/apiserver/apis/audit,sdn.io/sdwan/pkg/apiserver/apis/audit/v1beta1,sdn.io/sdwan/pkg/apiserver/apis/example,sdn.io/sdwan/pkg/apiserver/apis/example/v1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ifconfig,sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/zoneset,sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/vpn,sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/staticroute,sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/wifi,sdn.io/sdwan/pkg/apis/cubs/wifi/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/webauth,sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alert,sdn.io/sdwan/pkg/apis/cubs/alert/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/firewall,sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/businesspolicy,sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/organization,sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corporation,sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/profile,sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/site,sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/services,sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/clientvpn,sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/corptocorp,sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/guiuser,sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/check,sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/serviceclass,sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/troubleshoot,sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/devicedump,sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd,sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/orchpath,sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/alg,sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/multicastroute,sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/ztp,sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/blackfw,sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1"
# openapi
cd ../openapi-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1,sdn.io/sdwan/pkg/apimachinery/apis/meta/v1,sdn.io/sdwan/pkg/apimachinery/apis/meta/v1alpha1,sdn.io/sdwan/pkg/apimachinery/runtime,sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/wifi/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/alert/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1" \
               --output-package="sdn.io/sdwan/pkg/apigenerated/cubs/openapi"

# clientset
cd ../client-gen
go run main.go --input="batch/v1" --input-base="sdn.io/sdwan/pkg/apiserver/apis" --clientset-path="sdn.io/sdwan/pkg/apiclient"
go run main.go --input="runcmd/v1alpha1,devicedump/v1alpha1,firewall/v1alpha1,vpn/v1alpha1,ifconfig/v1alpha1,zoneset/v1alpha1,organization/v1alpha1,corporation/v1alpha1,profile/v1alpha1,site/v1alpha1,services/v1alpha1,businesspolicy/v1alpha1,staticroute/v1alpha1,webauth/v1alpha1,clientvpn/v1alpha1,corptocorp/v1alpha1,guiuser/v1alpha1,check/v1alpha1,serviceclass/v1alpha1,troubleshoot/v1alpha1,orchpath/v1alpha1,alg/v1alpha1,multicastroute/v1alpha1,ztp/v1alpha1,blackfw/v1alpha1" --input-base="sdn.io/sdwan/pkg/apis/cubs" --clientset-path="sdn.io/sdwan/pkg/apigenerated/cubs"


# listers
cd ../lister-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apiserver/apis/batch/v1" --output-package="sdn.io/sdwan/pkg/apiclient/listers"
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1" --output-package="sdn.io/sdwan/pkg/apigenerated/cubs/listers"

# informers
cd ../informer-gen
go run main.go --input-dirs="sdn.io/sdwan/pkg/apiserver/apis/batch/v1" --versioned-clientset-package="sdn.io/sdwan/pkg/apiclient/clientset" --listers-package="sdn.io/sdwan/pkg/apiclient/listers" --output-package="sdn.io/sdwan/pkg/apiclient/informers" --single-directory=true
go run main.go --input-dirs="sdn.io/sdwan/pkg/apis/cubs/runcmd/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/devicedump/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/firewall/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/vpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/ifconfig/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/zoneset/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/organization/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corporation/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/profile/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/site/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/services/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/businesspolicy/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/staticroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/webauth/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/clientvpn/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/corptocorp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/guiuser/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/check/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/serviceclass/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/troubleshoot/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/orchpath/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/alg/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/multicastroute/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/ztp/v1alpha1,sdn.io/sdwan/pkg/apis/cubs/blackfw/v1alpha1" --versioned-clientset-package="sdn.io/sdwan/pkg/apigenerated/cubs/clientset" --listers-package="sdn.io/sdwan/pkg/apigenerated/cubs/listers" --output-package="sdn.io/sdwan/pkg/apigenerated/cubs/informers" --single-directory=true
