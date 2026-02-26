# writesonic-cli

A command-line interface for the [Writesonic API](https://docs.writesonic.com/reference/introduction) — generate blog ideas, full articles, landing pages, marketing copy, and transform content directly from your terminal.

## Installation

```bash
go install github.com/the20100/writesonic-cli@latest
```

Or build from source:

```bash
git clone https://github.com/the20100/writesonic-cli
cd writesonic-cli
go build -o writesonic .
sudo mv writesonic /usr/local/bin/
```

## Authentication

Get your API key from [writesonic.com](https://app.writesonic.com/api) and save it:

```bash
writesonic auth set-key YOUR_API_KEY
```

Or use an environment variable:

```bash
export WRITESONIC_API_KEY=YOUR_API_KEY
```

## Quick Start

```bash
# Set your API key
writesonic auth set-key YOUR_API_KEY

# Configure defaults (optional)
writesonic auth config --engine premium --language en --copies 3

# Generate blog ideas
writesonic blog-ideas --topic "AI tools for marketers"

# Write an instant article
writesonic article instant --title "How to Use AI in Your Marketing Strategy"

# Generate landing page copy
writesonic landing page \
  --name "Acme SaaS" \
  --desc "Project management for remote teams" \
  --f1 "Task tracking" \
  --f2 "Team collaboration" \
  --f3 "Real-time analytics"
```

## Global Flags

Every command supports these flags:

| Flag | Default | Description |
|------|---------|-------------|
| `--engine` | `good` | AI quality: `economy`, `average`, `good`, `premium` |
| `--lang` | `en` | Language code (`fr`, `de`, `es`, `ja`, `zh`, etc.) |
| `--copies` | `1` | Number of variations to generate (1–5) |
| `--json` | | Force JSON output |
| `--pretty` | | Pretty-printed JSON |

## Commands

### `auth` — Authentication & Configuration

```bash
writesonic auth set-key YOUR_API_KEY        # Save API key
writesonic auth status                       # Show key and defaults
writesonic auth config --engine premium      # Set persistent defaults
writesonic auth logout                       # Remove stored key
```

### `blog-ideas` — Blog Post Ideas

Generate blog title ideas for a topic.

```bash
writesonic blog-ideas --topic "sustainable fashion"
writesonic blog-ideas --topic "AI in 2025" --copies 5 --keyword "machine learning"
```

### `article` — Full Articles

```bash
# Instant 1500-word article from a title
writesonic article instant --title "How to Learn Python in 30 Days"

# Long-form SEO article with sections
writesonic article write \
  --title "10 AI Tools Transforming Marketing" \
  --intro "AI is reshaping how marketers work..." \
  --sections "Writing Tools,Image Generation,Analytics,Future"
```

### `landing` — Landing Page Copy

```bash
# Full landing page (title, features, CTA, button)
writesonic landing page \
  --name "MyApp" --desc "Description" \
  --f1 "Feature 1" --f2 "Feature 2" --f3 "Feature 3"

# Headlines only
writesonic landing headline --name "MyApp" --desc "Description" --copies 5
```

### `copy` — Marketing Copy Frameworks

```bash
writesonic copy pas   --name "MyProduct" --desc "Description"   # Pain-Agitate-Solution
writesonic copy aida  --name "MyProduct" --desc "Description"   # Attention-Interest-Desire-Action
writesonic copy cta   --name "MyProduct" --copies 5              # Calls to Action
writesonic copy bullets --question "What are the benefits of X?" # Bullet-point answers
```

### `rewrite` — Content Transformation

```bash
writesonic rewrite rephrase  --content "Text to rephrase" --tone "formal"
writesonic rewrite shorten   --content "Long text to condense" --tone "casual"
writesonic rewrite tone      --content "Hey buddy!" --tone "professional"
writesonic rewrite keywords  --content "We build software." --keywords "SaaS, automation"
```

### `write` — Specific Content Pieces

```bash
writesonic write paragraph  --topic "Benefits of remote work" --instructions "Focus on productivity"
writesonic write meta       --title "My Blog Post" --desc "About AI tools"
writesonic write conclusion --topic "The future of AI in content creation"
```

### `update` — Self-update

Pull the latest source from GitHub, rebuild, and replace the current binary.

```bash
writesonic update
```

Requires `git` and `go` to be installed.

---

## Output Modes

Output is auto-detected:
- **Terminal** → human-readable text or tables
- **Piped** → JSON automatically

Override with `--json` or `--pretty`:

```bash
# Pretty JSON
writesonic blog-ideas --topic "AI" --pretty

# Pipe to jq
writesonic blog-ideas --topic "AI" --copies 5 | jq '.[].text'

# Save to file
writesonic article instant --title "My Article" --pretty > article.json
```

## Engines

| Engine | Speed | Quality | Use Case |
|--------|-------|---------|----------|
| `economy` | Fastest | Basic | Drafts, bulk generation |
| `average` | Fast | Good | Everyday content |
| `good` | Moderate | Very good | Default, most use cases |
| `premium` | Slower | Best | Final copy, important content |

## Supported Languages

25+ languages: `en`, `fr`, `de`, `es`, `it`, `pt-br`, `pt-pt`, `nl`, `pl`, `ru`, `ja`, `zh`, `sv`, `da`, `fi`, `el`, `hu`, `ro`, `cs`, `sk`, `sl`, `bg`, `lt`, `lv`, `et`

## Configuration File

Defaults are stored at:
- **macOS**: `~/Library/Application Support/writesonic/config.json`
- **Linux**: `~/.config/writesonic/config.json`
- **Windows**: `%AppData%\writesonic\config.json`

```json
{
  "api_key": "YOUR_KEY",
  "default_engine": "premium",
  "default_language": "fr",
  "default_copies": 3
}
```

## License

MIT
