library 'magic-butler-catalogue'
def PROJECT_NAME = 'terraform-provider-mezmo'
def DEFAULT_BRANCH = 'main'
def CURRENT_BRANCH = [env.CHANGE_BRANCH, env.BRANCH_NAME]?.find{branch -> branch != null}

def CREDS = [
    aws(
      credentialsId: 'aws',
      accessKeyVariable: 'AWS_ACCESS_KEY_ID',
      secretKeyVariable: 'AWS_SECRET_ACCESS_KEY'
    ),
    string(
      credentialsId: 'github-api-token',
      variable: 'GITHUB_TOKEN'
    )
]

def slugify(str) {
  def s = str.toLowerCase()
  s = s.replaceAll(/[^a-z0-9\s-\/]/, "").replaceAll(/\s+/, " ").trim()
  s = s.replaceAll(/[\/\s]/, '-').replaceAll(/-{2,}/, '-')
  s
}

pipeline {
  agent {
    node {
      label 'ec2-fleet'
      customWorkspace("/tmp/workspace/${env.BUILD_TAG}")
    }
  }

  parameters {
    string(name: 'SANITY_BUILD', defaultValue: '', description: 'Is this a scheduled sanity build that skips releasing?')
  }
  triggers {
    parameterizedCron(
      // Cron hours are in GMT, so this is roughly 12-3am EST, depending on DST
      env.BRANCH_NAME == DEFAULT_BRANCH ? 'H H(5-6) * * * % SANITY_BUILD=true' : ''
    )
  }

  options {
    timeout time: 1, unit: 'HOURS'
    timestamps()
    ansiColor 'xterm'
    withCredentials(CREDS)
  }

  environment {
    FEATURE_TAG = slugify("${CURRENT_BRANCH}-${BUILD_NUMBER}")
  }

  post {
    always {
      script {
        jiraSendBuildInfo()

        if (env.SANITY_BUILD == 'true') {
          notifySlack(
            currentBuild.currentResult,
            [channel: '#pipeline-bots'],
            "`${PROJECT_NAME}` sanity build took ${currentBuild.durationString.replaceFirst(' and counting', '')}."
          )
        }
      }
    }
  }

  stages {
    stage('Test') {
      when {
        not {
          changelog '\\[skip ci\\]'
        }
      }

      steps {
        sh 'printenv'
        sh 'echo "--starting"'
        sh 'make test'
      }
    }
  }
}
