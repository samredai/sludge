---
title: "Frequently Asked Questions"
tags: ["sludge", "faqs", "slurm", "hpc", "tview", "golang"]
aliases: ["frequently-asked-questions", "faq", "f", "q"]
ShowToc: true
TocOpen: true
---

# Why does sludge exist?

There are a few graphical interfaces out there for monitoring and interacting with a SLURM cluster such as [sview](https://slurm.schedmd.com/sview.html)
as well as [slurm-web](https://github.com/edf-hpc/slurm-web) which ships with SLURM. However, those require either a Remote Desktop Connection
or a web browser, or sometimes both depending on how access to your cluster is configured. Those who feel most comfortable in a terminal
shell often opt out of going through the trouble to connect to these UIs and instead rely on simply running commands through an SSH connection
to a head node (`sbatch`, `squeue`, `scancel`, etc.)

The goal of sludge is to provide a rich user interface that's directly accessible from the terminal while connected to a head node. The project is
inspired by other terminal UIs such as [lazydocker](https://github.com/jesseduffield/lazydocker) and [k9s](https://github.com/derailed/k9s).

Another motivating factor is that I just want to get familiar with building terminal UIs and found this as a good opportunity to learn. 😅

# What do I do if sludge is using up too many resources on a head node with many users?

Rendering and navigating the UI uses very little resources, however making repeated RPC calls to the SLURM cluster to keep the data on the screen fresh
can be resource intensive if the refresh interval is very small. By default, the refresh interval is set to `30s` and is usually passive enough as to not cause too much strain on the head node. However, you can change it using the `--refresh-interval` arg when launching sludge.

```sh
sludge --refresh-interval 1m
```

# How do I filter all of the data to just a single user, account, or both?

When launching sludge, you can use the `--username` or `--account` args to filter the data that's rendered to the screen.
They can also be used together to filter to a specific username and account which is useful when a user belongs to many accounts.

```sh
sludge --username samredai --account heavy_metal_lab
```