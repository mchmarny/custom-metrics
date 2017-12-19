import ConfigParser
from google.cloud import monitoring

Config = ConfigParser.ConfigParser()
Config.read("consume-metrics.cfg")

def ConfigSectionMap(section):
  configMap = {}
  options = Config.options(section)
  for option in options:
    try:
      configMap[option] = Config.get(section, option)
      if configMap[option] == -1:
        print("skip: %s" % option)
    except:
      print("exception on %s!" % option)
      configMap[option] = None
  return configMap

def get_service_account_auth(project, gcsJsonFile):
  client = monitoring.Client(project=project).from_service_account_json(gcsJsonFile)
  return client
  
def list_metrics(serviceAuthClient, metricName, metricDuration):
  query = serviceAuthClient.query(metricName, minutes=int(metricDuration))
  print(query.as_dataframe())
  
if __name__ == '__main__':
  gcsJsonFile = ConfigSectionMap("auth")['jsonfile']
  project = ConfigSectionMap("auth")['projectid']
  metric = ConfigSectionMap("metrics")['mockmetric']
  metricName = metric.split(':')[0]
  metricDuration = metric.split(':')[1]
  serviceAuthClient = get_service_account_auth(project, gcsJsonFile)
  list_metrics(serviceAuthClient, metricName, metricDuration)

