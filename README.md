# August CLI

Programmatic access to August. Authenticate every request with a Clerk API key: `Authorization: Bearer ak_...`.

Created by [@JustlyAI](https://github.com/JustlyAI) (justlyai).

## Install

The recommended path installs both the `august-pp-cli` binary and the `pp-august` agent skill (Claude Code, Codex, Cursor, Gemini CLI, GitHub Copilot, and other agents supported by the upstream [`skills`](https://github.com/vercel-labs/skills) CLI) in one shot:

```bash
npx -y @mvanhorn/printing-press-library install august
```

For CLI only (no skill):

```bash
npx -y @mvanhorn/printing-press-library install august --cli-only
```

For skill only — installs the skill into the same agents as the default command above, but skips the CLI binary (use this to update or reinstall just the skill):

```bash
npx -y @mvanhorn/printing-press-library install august --skill-only
```

To constrain the skill install to one or more specific agents (repeatable — agent names match the [`skills`](https://github.com/vercel-labs/skills) CLI):

```bash
npx -y @mvanhorn/printing-press-library install august --agent claude-code
npx -y @mvanhorn/printing-press-library install august --agent claude-code --agent codex
```

### Without Node

The generated install path is category-agnostic until this CLI is published. If `npx` is not available before publish, install Node or use the category-specific Go fallback from the public-library entry after publish.

### Pre-built binary

Download a pre-built binary for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/august-current). On macOS, clear the Gatekeeper quarantine: `xattr -d com.apple.quarantine <binary>`. On Unix, mark it executable: `chmod +x <binary>`.

<!-- pp-hermes-install-anchor -->
## Install for Hermes

Install the CLI binary first. The installer writes binaries to a per-user managed bin directory by default: `$HOME/.local/bin` on macOS/Linux and `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows.

```bash
npx -y @mvanhorn/printing-press-library install august --cli-only
```

Then install the focused Hermes skill.

From the Hermes CLI:

```bash
hermes skills install mvanhorn/printing-press-library/cli-skills/pp-august --force
```

Inside a Hermes chat session:

```bash
/skills install mvanhorn/printing-press-library/cli-skills/pp-august --force
```

Restart the Hermes session or gateway if the newly installed skill is not visible immediately.

## Install for OpenClaw
Install both the CLI binary and the focused OpenClaw skill. The installer defaults binaries to a per-user bin directory (`$HOME/.local/bin` on macOS/Linux, `%LOCALAPPDATA%\Programs\PrintingPress\bin` on Windows):

```bash
npx -y @mvanhorn/printing-press-library install august --agent openclaw
```

Restart the OpenClaw session or gateway if the newly installed skill is not visible immediately.

## Use with Claude Desktop

This CLI ships an [MCPB](https://github.com/modelcontextprotocol/mcpb) bundle — Claude Desktop's standard format for one-click MCP extension installs (no JSON config required).

To install:

1. Download the `.mcpb` for your platform from the [latest release](https://github.com/mvanhorn/printing-press-library/releases/tag/august-current).
2. Double-click the `.mcpb` file. Claude Desktop opens and walks you through the install.
3. Fill in `AUGUST_AUTHORIZATION` when Claude Desktop prompts you.

Requires Claude Desktop 1.0.0 or later. Pre-built bundles ship for macOS Apple Silicon (`darwin-arm64`) and Windows (`amd64`, `arm64`); for other platforms, use the manual config below.

<details>
<summary>Manual JSON config (advanced)</summary>

If you can't use the MCPB bundle (older Claude Desktop, unsupported platform), install the MCP binary and configure it manually.


Install the MCP binary from this CLI's published public-library entry or pre-built release.

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "august": {
      "command": "august-pp-mcp",
      "env": {
        "AUGUST_AUTHORIZATION": "<your-key>"
      }
    }
  }
}
```

</details>

## Quick Start

### 1. Install

See [Install](#install) above.

### 2. Set Up Credentials

Get your access token from your API provider's developer portal, then store it:

```bash
august-pp-cli auth set-token YOUR_TOKEN_HERE
```

Or set it via environment variable:

```bash
export AUGUST_AUTHORIZATION="your-token-here"
```

### 3. Verify Setup

```bash
august-pp-cli doctor
```

This checks your configuration and credentials.

### 4. Try Your First Command

```bash
august-pp-cli chats list
```

## Usage

Run `august-pp-cli --help` for the full command reference and flag list.

## Commands

### chats

Manage chats

- **`august-pp-cli chats by-id`** - Get a chat thread
- **`august-pp-cli chats cancel`** - Cancels a running or pending question. No-op if already in a terminal state. `question_id` is the chat_questions id returned by POST /chats.
- **`august-pp-cli chats create`** - Submits a query and returns the chat and question ids. The answer is produced asynchronously — poll `GET /chats/{chatId}` for the result.
- **`august-pp-cli chats delete`** - Permanently deletes a chat and all its messages (FK cascade). Caller must be the chat owner. Irreversible.
- **`august-pp-cli chats list`** - List chats
- **`august-pp-cli chats question-files`** - Lists work-product documents generated by a genius-mode question (genius_outputs folders).
- **`august-pp-cli chats question-result`** - Returns the answer, follow-up questions, and citation ids for a completed question. Citation excerpt resolution (doc name/text) is not yet available.
- **`august-pp-cli chats question-status`** - Polls the status of a submitted question, with a lightweight step-count progress signal.
- **`august-pp-cli chats rename`** - Renames a chat. Caller must be the chat owner.

### content_search

Manage content search

- **`august-pp-cli content-search`** - Full-text (BM25) search across the parsed text of documents the caller can access. Scope with folderIds and/or docIds; returns matching chunks with highlighted snippets.

### files

Manage files


### folders

Manage folders

- **`august-pp-cli folders create`** - Three positioning modes: subfolder (parentId), project root (projectId), or personal root (neither). Caller becomes owner.
- **`august-pp-cli folders delete`** - Soft-deletes one or more folders and their whole subtrees (owner-only). folderIds serialize as repeated query params.
- **`august-pp-cli folders list`** - Folders (alphabetical) then files (paginated by cursor/limit). Omit parentFolderId (null) for the personal/project root. All inputs are optional query params; arrays (filterTagIds) serialize as repeated query params. Returns items[] plus total_folders, total_files, nextCursor.

### pre-share-policies

Manage pre share policies

- **`august-pp-cli pre-share-policies pre-shares-create-policy`** - Create a pre-share policy
- **`august-pp-cli pre-share-policies pre-shares-delete-policy`** - Delete a pre-share policy
- **`august-pp-cli pre-share-policies pre-shares-get-policy`** - Get a pre-share policy
- **`august-pp-cli pre-share-policies pre-shares-list-policies`** - List pre-share policies
- **`august-pp-cli pre-share-policies pre-shares-toggle-policy-enabled`** - Enable or disable a pre-share policy

### projects

Manage projects

- **`august-pp-cli projects create`** - Creates a new project owned by the caller. The caller is added as the project owner.
- **`august-pp-cli projects delete`** - Permanently deletes a project and its resource links (chats and folders survive as personal resources). Destructive and irreversible; pass confirm=true.
- **`august-pp-cli projects list`** - Returns every project the authenticated user can access, most recently updated first.
- **`august-pp-cli projects rename`** - Renames a project. Caller must be the project owner.

### search_resource

Manage search resource

- **`august-pp-cli search-resource`** - Name search across the accessible folder tree. Pass projectId (optional query param) to scope to a project; omit for the caller's personal workspace. Returns a flat array of folder/file hits.

### workflows

Manage workflows

- **`august-pp-cli workflows`** - Generates a runnable workflow draft from a natural-language prompt and saves it. Reuse idempotencyKey to dedupe retries. Trigger provisioning is best-effort — failures return in triggerWarnings.


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
august-pp-cli chats list

# JSON for scripting and agents
august-pp-cli chats list --json

# Filter to specific fields
august-pp-cli chats list --json --select id,name,status

# Dry run — show the request without sending
august-pp-cli chats list --dry-run

# Agent mode — JSON + compact + no prompts in one flag
august-pp-cli chats list --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries and `--ignore-missing` to delete retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Health Check

```bash
august-pp-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/august-public-pp-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `AUGUST_AUTHORIZATION` | per_call | Yes | Set to your API credential. |

### agentcookie (optional)

If you use agentcookie to sync secrets across machines, this CLI auto-adopts agentcookie-managed credentials with no extra setup. When the daemon writes to this CLI's config, `august-pp-cli doctor` reports `agentcookie: detected` and `auth-status` labels the source as `agentcookie`. Skip this section if you don't use agentcookie - the CLI works the same as any other.

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `august-pp-cli doctor` to check credentials
- Verify the environment variable is set: `echo $AUGUST_AUTHORIZATION`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

---

Generated by [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press)
