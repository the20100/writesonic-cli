package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vincentmaurin/writesonic-cli/internal/output"
)

var (
	articleTitle    string
	articleIntro    string
	articleSections string
)

var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "Generate full articles",
}

var articleV3Cmd = &cobra.Command{
	Use:   "write",
	Short: "Generate a long-form SEO article (AI Article Writer v3)",
	Example: `  writesonic article write --title "10 AI Tools in 2025" --intro "AI is transforming..." --sections "Tools,Use cases,Future"
  writesonic article write --title "Healthy Eating" --intro "Good nutrition is key" --sections "Benefits,Tips,Recipes" --copies 1`,
	RunE: runArticleV3,
}

var articleInstantCmd = &cobra.Command{
	Use:   "instant",
	Short: "Generate a 1500-word article instantly",
	Example: `  writesonic article instant --title "How to Learn Python in 2025"
  writesonic article instant --title "Best Coffee Shops in Paris" --engine premium`,
	RunE: runArticleInstant,
}

var (
	instantTitle string
)

func init() {
	articleV3Cmd.Flags().StringVar(&articleTitle, "title", "", "Article title (required)")
	articleV3Cmd.Flags().StringVar(&articleIntro, "intro", "", "Article introduction (required)")
	articleV3Cmd.Flags().StringVar(&articleSections, "sections", "", "Comma-separated section titles (required)")
	articleV3Cmd.MarkFlagRequired("title")
	articleV3Cmd.MarkFlagRequired("intro")
	articleV3Cmd.MarkFlagRequired("sections")

	articleInstantCmd.Flags().StringVar(&instantTitle, "title", "", "Article title (required)")
	articleInstantCmd.MarkFlagRequired("title")

	articleCmd.AddCommand(articleV3Cmd, articleInstantCmd)
	rootCmd.AddCommand(articleCmd)
}

func runArticleV3(cmd *cobra.Command, args []string) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	// Parse comma-separated sections into slice
	rawSections := strings.Split(articleSections, ",")
	sections := make([]string, 0, len(rawSections))
	for _, s := range rawSections {
		s = strings.TrimSpace(s)
		if s != "" {
			sections = append(sections, s)
		}
	}

	body := map[string]interface{}{
		"article_title":    articleTitle,
		"article_intro":    articleIntro,
		"article_sections": sections,
	}

	results, err := client.PostResults("/ai-article-writer-v3", params, body)
	if err != nil {
		return err
	}

	if output.IsJSON(jsonFlag, prettyFlag) {
		return output.PrintJSON(results, prettyFlag)
	}

	texts := make([]string, len(results))
	for i, r := range results {
		texts[i] = r.Text
	}
	output.PrintText(texts)
	return nil
}

func runArticleInstant(cmd *cobra.Command, args []string) error {
	params := url.Values{}
	params.Set("engine", engineFlag)
	params.Set("language", langFlag)
	params.Set("num_copies", fmt.Sprintf("%d", copiesFlag))

	body := map[string]interface{}{
		"article_title": instantTitle,
	}

	results, err := client.PostResults("/instant-article-writer", params, body)
	if err != nil {
		return err
	}

	if output.IsJSON(jsonFlag, prettyFlag) {
		return output.PrintJSON(results, prettyFlag)
	}

	texts := make([]string, len(results))
	for i, r := range results {
		texts[i] = r.Text
	}
	output.PrintText(texts)
	return nil
}
