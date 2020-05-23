import jenkins.model.Jenkins
import org.csanchez.jenkins.plugins.kubernetes.KubernetesCloud

def kubernetes = new KubernetesCloud("kubernetes")
kubernetes.setMaxRequestsPerHost(KubernetesCloud.DEFAULT_MAX_REQUESTS_PER_HOST)
kubernetes.setDirectConnection(true)

Jenkins.instance.clouds.removeAll(KubernetesCloud)
Jenkins.instance.clouds.add(kubernetes)