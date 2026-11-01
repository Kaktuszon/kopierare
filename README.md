# Kopierare
Small utility program written in Go to easy copy things to clipboard that you might want often.\

Supports copying with help of xclip and wl-copy.

config.yaml example:
```yaml
software: xclip
lines:
  - Anka
  - Banan
  - Senap
```

config.yaml should be placed in ~/.kopierare/config.yaml or specified with launch argument -path.

__This could be a Bash script :D__