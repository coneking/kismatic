package cli

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

const (
	defaultDBName        = filepath.Join(assetsFolder, "clusterStates.db")
	defaultPlanName      = "kismatic-cluster.yaml"
	defaultClusterName   = "kubernetes"
	defaultGeneratedName = "generated"
	defaultRunsName      = "runs"
	defaultTimeout       = 10 * time.Second
	clustersBucket       = "kismatic"
	assetsFolder         = "clusters"
	defaultInsecurePort  = "8080"
	defaultSecurePort    = "8443"
)

type planFileNotFoundErr struct {
	filename string
}

func (e planFileNotFoundErr) Error() string {
	return fmt.Sprintf("Plan file not found at %q. If you don't have a plan file, you may generate one with 'kismatic install plan'", e.filename)
}

// Returns a path to a plan file, generated dir, and runs dir according to the clusterName
func generateDirsFromName(clusterName string) (string, string, string) {
	return filepath.Join(assetsFolder, clusterName, defaultPlanName), filepath.Join(assetsFolder, clusterName, defaultGeneratedName), filepath.Join(assetsFolder, clusterName, defaultRunsName)
}

// CheckClusterExists does a simple check to see if the cluster folder+plan file exists in clusters
func CheckClusterExists(name string) (bool, error) {
	files, err := ioutil.ReadDir(assetsFolder)
	if err != nil {
		return false, err
	}
	for _, finfo := range files {
		if finfo.Name() == name {
			possiblePlans, err := ioutil.ReadDir(filepath.Join(assetsFolder, finfo.Name()))
			if err != nil {
				return false, err
			}
			for _, possiblePlan := range possiblePlans {
				if possiblePlan.Name() == defaultPlanName {
					return true, nil
				}
			}
		}
	}
	return false, fmt.Errorf("Cluster with name %s not found. If you have a plan file, but your cluster doesn't exist, please run kismatic import PLAN_FILE_PATH.", name)
}

// CheckPlaybookExists does a check to make sure the step exists
func CheckPlaybookExists(play string) (bool, error) {
	plays, err := ioutil.ReadDir("ansible/playbooks")
	if err != nil {
		return false, err
	}
	for _, finfo := range plays {
		if finfo.Name() == play {
			return true, nil
		}
	}
	return false, fmt.Errorf("playbook %s not found")
}

//
func CreateStoreIfNotExists(path string) error {
	dbPath = path
	if len(path) == 0 {
		dbPath = defaultDBName
	}
}
