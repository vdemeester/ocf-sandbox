package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	projectv1 "github.com/openshift/api/project/v1"
	projectapiv1 "github.com/openshift/origin/pkg/project/apis/project/v1"
)

func init() {
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(projectapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(projectv1.GroupVersion))
}
