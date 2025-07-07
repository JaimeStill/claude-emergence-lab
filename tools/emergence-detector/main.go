package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// Define command-line flags
	var (
		jsonOutput = flag.Bool("json", false, "Output results as JSON")
		verbose    = flag.Bool("v", false, "Verbose output")
		threshold  = flag.Float64("threshold", 0.7, "Confidence threshold (0.0-1.0)")
		help       = flag.Bool("help", false, "Show help message")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "emergence-detector - Identify emergence patterns in text\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nPatterns detected:\n")
		fmt.Fprintf(os.Stderr, "  - Symbiotic Combination: Elements creating more than their sum\n")
		fmt.Fprintf(os.Stderr, "  - Recursive Loop: Self-modifying cycles\n")
		fmt.Fprintf(os.Stderr, "  - Emergent Causation: Enabling conditions rather than direct cause\n")
	}
	
	flag.Parse()
	
	if *help || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(0)
	}
	
	filename := flag.Arg(0)
	detector := NewDetector(*verbose)
	
	// Analyze the file
	result, err := detector.AnalyzeFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Output results
	if *jsonOutput {
		outputJSON(result)
	} else {
		outputText(result, *threshold)
	}
}

func outputJSON(result *AnalysisResult) {
	output := map[string]interface{}{
		"source":     result.Source,
		"summary":    result.Summary,
		"detections": result.Detections,
	}
	
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}

func outputText(result *AnalysisResult, threshold float64) {
	fmt.Printf("Emergence Analysis: %s\n", result.Source)
	fmt.Printf("=====================================\n\n")
	
	if !result.HasEmergence() {
		fmt.Println("No emergence patterns detected.")
		return
	}
	
	// Summary
	fmt.Println("Summary:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for pattern, count := range result.Summary {
		if count > 0 {
			fmt.Fprintf(w, "  %s:\t%d instances\n", pattern, count)
		}
	}
	w.Flush()
	fmt.Println()
	
	// High confidence detections
	highConf := result.GetHighConfidenceDetections(threshold)
	if len(highConf) > 0 {
		fmt.Printf("High Confidence Detections (>= %.1f):\n", threshold)
		fmt.Println("=====================================")
		
		for i, detection := range highConf {
			fmt.Printf("\n%d. %s (confidence: %.2f)\n", i+1, detection.Pattern, detection.Confidence)
			fmt.Printf("   Location: %s\n", detection.Location)
			fmt.Printf("   Context: %s\n", truncate(detection.Context, 80))
			fmt.Printf("   Analysis: %s\n", detection.Explanation)
		}
	}
	
	// Overall assessment
	fmt.Println("\nEmergence Assessment:")
	fmt.Println("====================")
	assessEmergence(result)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func assessEmergence(result *AnalysisResult) {
	total := len(result.Detections)
	highConf := len(result.GetHighConfidenceDetections(0.8))
	
	if highConf > 5 {
		fmt.Println("✦ Strong emergence patterns detected")
		fmt.Println("  Multiple high-confidence patterns suggest complex emergent behavior")
	} else if highConf > 0 {
		fmt.Println("◐ Moderate emergence patterns detected")
		fmt.Println("  Some clear patterns of emergence are present")
	} else if total > 0 {
		fmt.Println("○ Potential emergence patterns detected")
		fmt.Println("  Lower confidence indicators suggest possible emergent behavior")
	}
	
	// Pattern-specific insights
	if result.Summary["Recursive Loop"] > 2 {
		fmt.Println("  ⟳ Multiple recursive patterns suggest self-organizing behavior")
	}
	if result.Summary["Symbiotic Combination"] > 2 {
		fmt.Println("  ⊕ Symbiotic combinations indicate synergistic effects")
	}
	if result.Summary["Emergent Causation"] > 2 {
		fmt.Println("  ⟹ Emergent causation suggests enabling conditions present")
	}
}