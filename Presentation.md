
Slide 1:
Hello, I'm Kaveesh Dubey. I work for the Azure Monitor team and my team focuses on the kuberenetes monitoring scenarios.
We recently announced public preview for monitoring the linux nodes on an AKS cluster using prometheus and we're in the process of releasing this feature for windows nodes too!

Slide 2:

Windows usage has been increasing in the kuberenetes eco-system over the last couple of years and there is great demand both from internal and external customers for using prometheus to monitor. Right now, there is no readily available managed prometheus offering that natively supports windows nodes in AKS and we believe this is a great opportunity for Microsoft to lead the way in the windows space.

Slide 3:
We deploy via the tried and tested addon method that customers are familiar with. This can be done in a variety of ways using the Azure Portal, CLI, ARM, Bicep etc. Once the addohn is enabled you get monitoring by default for both your linux and windows nodes. It deploys a daemonset pod on all of the windows nodes which we use to scrape default monitoring prometheus metrics. As part of enablement we also provision a select group of recording rules and dashboards that we've taken from the Open Source Community (OSS) and integrated them with the Azure Managed Grafana service. Customers can even add in any custom prometheus metrics they're exposing from their applications runnin on the cluster and update the addon's configuration to scrape these metrics and view them in Grafana.
One other key differentiator from the normal prometheus model is that all the metrics that are scraped and ingested are sent to an Azure Monitor Workspace and this does not take up any resource from the cluster.
The addon support both Windows Server 2019 and Windows Server 2022 as by the end of this year the AKS team plans on making Windows Server 2022 the default nodepool in AKS. We're also working with the AKS team so that every AKS node by default also have windows exporter which exposes prometheus metrics similar to how they have a node exporter for their linux nodepools.

Slide 4:
Show Demo!


Github Page with links :
For this demo I'll show our regular development and release process with our CI/CD pipelines.
These clusters are monitoring development and production clusters. Whenever any change is merged into the main branch of our codebase an Azure DevOps pipeline kicks in. It builds and deploys a new container image to all the development clusters. We have alerts setup on them to identify if anything starts misbehaving or if any data flow is breaking.
