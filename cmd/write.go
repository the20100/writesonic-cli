package cmd

import (
	"github.com/spf13/cobra"
)

// write.go groups standalone writing utilities: paragraph, meta, conclusion

var (
	paragraphTopic        string
	paragraphInstructions string
	metaBlogTitle         string
	metaBlogDesc          string
	conclusionTopic       string
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write specific content pieces (paragraphs, meta tags, conclusions)",
}

var writeParagraphCmd = &cobra.Command{
	Use:   "paragraph",
	Short: "Write a structured, persuasive paragraph",
	Example: `  writesonic write paragraph --topic "Benefits of remote work"
  writesonic write paragraph --topic "Why use AI for writing" --instructions "Focus on speed and quality"`,
	RunE: runWriteParagraph,
}

var writeMetaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Generate SEO meta title and description for a blog post",
	Example: `  writesonic write meta --title "10 Tips for Remote Work" --desc "Advice for distributed teams"
  writesonic write meta --title "Best Coffee Shops in Paris" --desc "A guide to Paris cafes" --copies 3`,
	RunE: runWriteMeta,
}

var writeConclusionCmd = &cobra.Command{
	Use:   "conclusion",
	Short: "Write a compelling conclusion for an article",
	Example: `  writesonic write conclusion --topic "The future of AI in content creation"
  writesonic write conclusion --topic "Remote work benefits" --copies 2`,
	RunE: runWriteConclusion,
}

func init() {
	writeParagraphCmd.Flags().StringVar(&paragraphTopic, "topic", "", "Topic to write about (required)")
	writeParagraphCmd.Flags().StringVar(&paragraphInstructions, "instructions", "", "Additional instructions (optional)")
	writeParagraphCmd.MarkFlagRequired("topic")

	writeMetaCmd.Flags().StringVar(&metaBlogTitle, "title", "", "Blog post title (required)")
	writeMetaCmd.Flags().StringVar(&metaBlogDesc, "desc", "", "Blog post description (required)")
	writeMetaCmd.MarkFlagRequired("title")
	writeMetaCmd.MarkFlagRequired("desc")

	writeConclusionCmd.Flags().StringVar(&conclusionTopic, "topic", "", "Article topic to conclude (required)")
	writeConclusionCmd.MarkFlagRequired("topic")

	writeCmd.AddCommand(writeParagraphCmd, writeMetaCmd, writeConclusionCmd)
	rootCmd.AddCommand(writeCmd)
}

func runWriteParagraph(cmd *cobra.Command, args []string) error {
	body := map[string]interface{}{
		"topic": paragraphTopic,
	}
	if paragraphInstructions != "" {
		body["instructions"] = paragraphInstructions
	}
	return postAndPrint("/paragraph-writer", body)
}

func runWriteMeta(cmd *cobra.Command, args []string) error {
	return postAndPrint("/meta-blog", map[string]interface{}{
		"blog_title":       metaBlogTitle,
		"blog_description": metaBlogDesc,
	})
}

func runWriteConclusion(cmd *cobra.Command, args []string) error {
	return postAndPrint("/conclusion-writer", map[string]interface{}{
		"topic": conclusionTopic,
	})
}
