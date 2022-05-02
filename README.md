# dooray-action.post-task

## Usage
### Notify Releases
```
name: Notify releases

on:
  release:
    types: [published]

jobs:
  notify-relase:
    name: Notify releases
    runs-on: ubuntu-latest
    steps:
      - name: Checkout sources
        uses: actions/checkout@v3
        with:
          repository: wagstyle/dooray-action.post-task
          ref: v1.0.0

      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.18

      - name: Post dooray task
        uses: wagstyle/dooray-action.post-task@v1.0.0
        env:
          PROJECT_ID: ${{ secrets.DOORAY_PROJECT_ID }}
          AUTHORIZATION_TOKEN: ${{ secrets.DOORAY_AUTHORIZATION_TOKEN }}
          SUBJECT: Release ${{ github.ref_name }}
          CONTENT: ${{ github.event.release.body }}
          RECIPIENT: '{"type": "group", "group": { "projectMemberGroupId": "3202204541800548226" }}'
          TAG: "3263696335409770355"

```
