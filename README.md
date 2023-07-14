# cloud-concierge
<p align="center">
<img width="500" src=./images/cloud-concierge-logo.png>
</p>
<h2 align="center">
<a href="https://docs.cloudconcierge.io" target="_blank">Docs</a> |
<a href="https://https://www.youtube.com/watch?v=y8OSfQQMEL0&t=12s" target="_blank"> Managed Demo </a>
</h2>

## Motivation
Many teams build their own Terraform management "stacks" using major cloud provider state backends
and tools like Atlantis for running `plan` and `apply` and state-locking. 

For more sophisticated tooling, some may turn to tools like Terraform Cloud,
Scalr, Spacelift and Firefly. We find, however, that these tool's pricing can become particularly onerous
when wanting to self-host runners or access the most desired features like drift detection, security scanning, etc.

## Why Cloud Concierge?
Cloud Concierge is an open-sourced container that integrates with your existing Terraform management stack.
All results and codified resources are output in a digestible Pull Request to a repository of your choice, providing you with a "State of Cloud"
report in a GitOps manner. It provides:
- Cloud codification, identify un-managed resources and generate corresponding Terraform code and import statements/import blocks
<p align="center">
<img width="650" src=./images/tool-screenshots/generated_resources.png alt="generated_resources">
</p>

- Drift detection
<p align="center">
<img width="650" src=./images/tool-screenshots/drift_detection.png alt="drift_detection">
</p>

- Flag accounts creating changes outside your Terraform workflow
<p align="center">
<img width="650" src=./images/tool-screenshots/cloud_actors.png alt="cloud_actor_identification">
</p>

- Whole-cloud cost estimation, powered by Infracost
<p align="center">
<img width="650" src=./images/tool-screenshots/costs_I.png alt="cost_estimations_aggregate">
</p>

<p align="center">
<img width="650" src=./images/tool-screenshots/costs_II.png alt="cost_estimations_individual">
</p>

- Whole-cloud security scanning, powered by tfsec (checkov integration coming soon)
<p align="center">
<img width="650" src=./images/tool-screenshots/security_risk_id.png alt="security_risks">
</p>

## How does it work?
1) Cloud Concierge creates a representation of your cloud infrastructure as Terraform
2) This representation is compared against your state files to detect drift, and identify resources outside of Terraform control
3) Static security scans and cost estimation is performed on the Terraform representation
4) Results and code are summarized in a [Pull Request](https://docs.cloudconcierge.io/how-it-works/pull-request-output) within the repository of your choice

## Getting Started
0) Retrieve an organization token from the dragondrop.cloud management platform [here](https://app.dragondrop.cloud).
1) Configure your environment variable file. This determines the execution behavior of the container. We provide example env configuration files for:
   - [AWS](./examples/environments/aws-example.env)
   - [GCP](./examples/environments/gcp-example.env)
   - [Azure](./examples/environments/azure-example.env)

Detailed documentation on environment variables needed can be found [here](https://docs.cloudconcierge.io/running-cloud-concierge/environment-variables).

While Cloud Concierge validates environment variable formats upon start-up, we provide a UI for client-side validation of env vars
within the [dragondrop.cloud platform](https://app.dragondrop.cloud/env-var-validator) should faster iteration be desired.

4) Run the container with the following command:
```bash
docker run --env-file ./path/to/my/env-file.env -v main:/main -w /main  dragondropcloud/cloud-concierge:latest
```

5) If using Terraform >= 1.5, Cloud Concierge generates [import blocks](https://medium.com/@hello_9187/terraform-1-5-xs-new-import-block-b8607c51287f) for newly codified resources directly.
If using Terraform < 1.5, we generate a `terraform import` command for each resource. These commands can be run manually,
or programmatically in a `plan` and `apply` manner using our [GitHub Action](https://github.com/dragondrop-cloud/github-action-tfstate-migration). 

### Running on a schedule
A common use case is to want to regularly scan for drift and un-codified resources. Cloud Concierge can easily be run
on a cron schedule using GitHub Actions. See our [example workflow](https://github.com/dragondrop-cloud/cloud-concierge/blob/dev/examples/github_action.yml).

### Telemetry
For OSS usage, Cloud Concierge only logs whenever a container execution is started. This method can be viewed [here](pkg/implementations/dragon_drop/http_dragondrop_oss_methods.go).
 
Jobs managed by the [dragondrop platform](https://dragondrop.cloud) log statuses over the course of the job execution and anonymized data for cloud visualizations. These methods
can be viewed [here](pkg/implementations/dragon_drop/http_dragondrop_managed_execution.go) and [here]([here](pkg/implementations/dragon_drop/http_dragondrop_managed_execution.go)).

## Our Roadmap
We are just getting started, and have a lot of exciting features on our roadmap. More details can be found [here](https://github.com/dragondrop-cloud/cloud-concierge/wiki/Roadmap).

## Contributing
Contributions in any form are highly encouraged. Check out our [contributing guide](CONTRIBUTING.md) to get started.

## Using at Scale w/dragondrop.cloud
The cloud-concierge container is easy to manage in a single configuration.
If you are looking to use cloud-concierge at scale, however, the [dragondrop.cloud](https://dragondrop.cloud/how-it-works) management platform allows you to:
- Manage multiple cloud-concierge configurations through a user interface
- Manage different cron jobs for executing each configuration at desired intervals
- Consolidate visibility across all cloud-concierge executions into visualizations of drift, uncodified resources, cloud costs, and security risks.
- Continue to self-host cloud-concierge instances within your cloud using [serverless infrastructure](https://registry.terraform.io/namespaces/dragondrop-cloud).

## Other Resources
- [Documentation](https://docs.cloudconcierge.io)
- [Slack](https://cloud-concierge.slack.com/join/shared_invite/zt-1xx3sqsb6-cekIXs2whccZvbU81Xn5qg#/shared-invite/email)
- [Tool Walk Through + Use Case (low stakes and no pressure!)](https://calendly.com/dragondrop-cloud/cloud-concierge-walk-through)
- [Terraform Learning Resources](https://dragondrop.cloud/learn/terraform/)
- [Medium Blog](https://medium.com/@hello_9187)
- [Managed Offering](https://dragondrop.cloud/how-it-works/)
