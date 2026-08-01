# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v1.3.1 — 2026-08-01

Infrastructure only. No library code changed — the package API and behaviour
are identical to v1.3.0.

- Mirroring and archiving live in their own `mirror-and-archive.yml` workflow,
  separate from the build `pipeline.yml`.
- Every branch and tag push is mirrored to Codeberg as well as GitLab. The
  mirrors are force-pushed from GitHub, so pull/merge requests are turned off
  on them — anything merged there would be destroyed by the next sync. Issues
  and forking stay enabled.
- Pushes to the default branch and to tags, plus a monthly cron, archive the
  repo to the Wayback Machine (through the authenticated Save Page Now API,
  outlinks included), Software Heritage, and archive.org.
- A new `issue-pull.yml` workflow runs every six hours and copies issues opened
  on either mirror back into GitHub, closing the GitHub copy when the original
  closes. The schedule is staggered per repo and the scheduled run adds jitter;
  a manual `workflow_dispatch` run does not wait.
- `.gitignore` now excludes `git-update.sh`, the per-release tooling script.
