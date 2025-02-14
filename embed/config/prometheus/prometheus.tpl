global:
  scrape_interval: 15s
  evaluation_interval: 15s

  external_labels:
    monitor: "hstream-monitor"

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          {{- range .AlertManagerAddress }}
          - '{{.}}'
          {{- end }}

rule_files:
  - "disks.yml"
  - "cluster.yml"
  - "zk.yml"
  - "alert.yml"

scrape_configs:
  - job_name: "monitor_port_probe"
    scrape_interval: 30s
    static_configs:
    - targets:
      {{- range .NodeExporterAddress }}
      - '{{.}}'
      {{- end }}
      labels:
        group: 'node_exporter'
    - targets:
      {{- range .CadVisorAddress }}
      - '{{.}}'
      {{- end }}
      labels:
        group: 'cadvisor'
    relabel_configs:
      - source_labels: [__address__]
        target_label: instance
        separator: ':'
        regex: '(.*):.*'
        replacement: "${1}"
  - job_name: "hstream_metrics"
    scrape_interval: 5s
    static_configs:
    - targets:
      {{- range .HStreamExporterAddress }}
      - '{{.}}'
      {{- end }}