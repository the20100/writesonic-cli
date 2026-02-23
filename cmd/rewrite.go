package cmd

import (
	"github.com/spf13/cobra"
)

// rewrite.go groups content transformation commands: rephrase, shorten, tone-changer, rewrite-with-keywords

var (
	rephraseContent  string
	rephraseTone     string
	shortenContent   string
	shortenTone      string
	toneContent      string
	toneTone         string
	kwContent        string
	kwKeywords       string
)

var rewriteCmd = &cobra.Command{
	Use:   "rewrite",
	Short: "Transform existing content (rephrase, shorten, tone, keywords)",
}

var rewriteRephraseCmd = &cobra.Command{
	Use:   "rephrase",
	Short: "Rephrase content in a different style",
	Example: `  writesonic rewrite rephrase --content "The quick brown fox jumps over the lazy dog."
  writesonic rewrite rephrase --content "Our product is amazing." --tone "formal"`,
	RunE: runRephrase,
}

var rewriteShortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Shorten content while keeping the message",
	Example: `  writesonic rewrite shorten --content "Our product is the most amazing and revolutionary tool on the market today."
  writesonic rewrite shorten --content "Long paragraph here..." --tone "casual"`,
	RunE: runShorten,
}

var rewriteToneCmd = &cobra.Command{
	Use:   "tone",
	Short: "Change the tone of existing content",
	Example: `  writesonic rewrite tone --content "Hey there! Check out our new product!" --tone "formal"
  writesonic rewrite tone --content "Our quarterly results show..." --tone "casual"`,
	RunE: runToneChanger,
}

var rewriteKeywordsCmd = &cobra.Command{
	Use:   "keywords",
	Short: "Rewrite content with target SEO keywords",
	Example: `  writesonic rewrite keywords --content "We sell software." --keywords "project management, team collaboration"
  writesonic rewrite keywords --content "Article text here" --keywords "AI, machine learning, automation"`,
	RunE: runRewriteKeywords,
}

func init() {
	rewriteRephraseCmd.Flags().StringVar(&rephraseContent, "content", "", "Content to rephrase (required, 20-1000 chars)")
	rewriteRephraseCmd.Flags().StringVar(&rephraseTone, "tone", "", "Desired tone of voice (optional)")
	rewriteRephraseCmd.MarkFlagRequired("content")

	rewriteShortenCmd.Flags().StringVar(&shortenContent, "content", "", "Content to shorten (required, 20-1000 chars)")
	rewriteShortenCmd.Flags().StringVar(&shortenTone, "tone", "", "Desired tone of voice (optional)")
	rewriteShortenCmd.MarkFlagRequired("content")

	rewriteToneCmd.Flags().StringVar(&toneContent, "content", "", "Content to transform (required)")
	rewriteToneCmd.Flags().StringVar(&toneTone, "tone", "", "Target tone (e.g. formal, casual, professional)")
	rewriteToneCmd.MarkFlagRequired("content")
	rewriteToneCmd.MarkFlagRequired("tone")

	rewriteKeywordsCmd.Flags().StringVar(&kwContent, "content", "", "Content to rewrite (required)")
	rewriteKeywordsCmd.Flags().StringVar(&kwKeywords, "keywords", "", "Comma-separated target keywords (required)")
	rewriteKeywordsCmd.MarkFlagRequired("content")
	rewriteKeywordsCmd.MarkFlagRequired("keywords")

	rewriteCmd.AddCommand(rewriteRephraseCmd, rewriteShortenCmd, rewriteToneCmd, rewriteKeywordsCmd)
	rootCmd.AddCommand(rewriteCmd)
}

func runRephrase(cmd *cobra.Command, args []string) error {
	body := map[string]interface{}{
		"content_to_rephrase": rephraseContent,
	}
	if rephraseTone != "" {
		body["tone_of_voice"] = rephraseTone
	}
	return postAndPrint("/content-rephrase", body)
}

func runShorten(cmd *cobra.Command, args []string) error {
	body := map[string]interface{}{
		"content_to_shorten": shortenContent,
	}
	if shortenTone != "" {
		body["tone_of_voice"] = shortenTone
	}
	return postAndPrint("/content-shorten", body)
}

func runToneChanger(cmd *cobra.Command, args []string) error {
	return postAndPrint("/tone-changer", map[string]interface{}{
		"content_to_change": toneContent,
		"tone":              toneTone,
	})
}

func runRewriteKeywords(cmd *cobra.Command, args []string) error {
	return postAndPrint("/rewrite-with-keywords", map[string]interface{}{
		"content":  kwContent,
		"keywords": kwKeywords,
	})
}
