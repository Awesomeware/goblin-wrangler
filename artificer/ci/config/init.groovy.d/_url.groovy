import jenkins.model.JenkinsLocationConfiguration

def jlc = JenkinsLocationConfiguration.get()
jlc.setUrl("https://ci.goblinwrangler.com")
jlc.save()