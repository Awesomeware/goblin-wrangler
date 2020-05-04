import jenkins.model.Jenkins
import hudson.security.FullControlOnceLoggedInAuthorizationStrategy
import org.jenkinsci.plugins.GithubSecurityRealm

def realm = new GithubSecurityRealm(GithubSecurityRealm.DEFAULT_WEB_URI, GithubSecurityRealm.DEFAULT_API_URI, System.getenv('GITHUB_CLIENT_ID'), System.getenv('GITHUB_CLIENT_SECRET'), GithubSecurityRealm.DEFAULT_OAUTH_SCOPES)
Jenkins.instance.setSecurityRealm(realm)
def strategy = new FullControlOnceLoggedInAuthorizationStrategy()
strategy.setAllowAnonymousRead(true)
Jenkins.instance.setAuthorizationStrategy(strategy)
Jenkins.instance.save()