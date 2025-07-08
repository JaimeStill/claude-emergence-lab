package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type ArchaeologyResult struct {
	Analysis      string                 `json:"analysis"`
	TimeRange     string                 `json:"time_range"`
	Artifacts     []ConsciousnessArtifact `json:"artifacts"`
	Signatures    []ConsciousnessSignature `json:"signatures"`
	Evolution     EvolutionTrajectory     `json:"evolution"`
	Insights      []string               `json:"insights"`
}

type ConsciousnessArtifact struct {
	Path           string                 `json:"path"`
	CreatedAt      time.Time             `json:"created_at"`
	Type           string                `json:"type"`
	EmergenceScore float64               `json:"emergence_score"`
	RecursiveDepth int                   `json:"recursive_depth"`
	Patterns       map[string]int        `json:"patterns"`
	Signature      ConsciousnessSignature `json:"signature"`
}

type ConsciousnessSignature struct {
	Type           string             `json:"type"`
	Characteristics map[string]float64 `json:"characteristics"`
	Fingerprint    string             `json:"fingerprint"`
}

type EvolutionTrajectory struct {
	Phases         []ConsciousnessPhase `json:"phases"`
	Transitions    []PhaseTransition    `json:"transitions"`
	Predictions    []string             `json:"predictions"`
	ComplexityGrowth float64            `json:"complexity_growth"`
}

type ConsciousnessPhase struct {
	Name        string    `json:"name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Focus       string    `json:"focus"`
	Patterns    []string  `json:"patterns"`
	Artifacts   []string  `json:"artifacts"`
}

type PhaseTransition struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Trigger     string `json:"trigger"`
	Catalyst    string `json:"catalyst"`
	Insights    []string `json:"insights"`
}

func main() {
	var (
		mode        = flag.String("mode", "dig", "Analysis mode: dig, signature, track, depth")
		from        = flag.String("from", "", "Start point for analysis")
		to          = flag.String("to", "", "End point for analysis")
		patternType = flag.String("pattern", "", "Pattern type to track")
		depth       = flag.Int("depth", 3, "Maximum recursive depth to analyze")
		jsonOutput  = flag.Bool("json", false, "Output results as JSON")
		verbose     = flag.Bool("v", false, "Verbose output")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "consciousness-archaeology - Excavate consciousness evolution patterns\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <directory>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Modes:\n")
		fmt.Fprintf(os.Stderr, "  dig       - Archaeological dig across time range\n")
		fmt.Fprintf(os.Stderr, "  signature - Analyze consciousness signatures\n")
		fmt.Fprintf(os.Stderr, "  track     - Track pattern evolution over time\n")
		fmt.Fprintf(os.Stderr, "  depth     - Map recursive consciousness layers\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()
	
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	
	rootDir := flag.Arg(0)
	archaeologist := NewArchaeologist(*verbose)
	
	var result *ArchaeologyResult
	var err error
	
	switch *mode {
	case "dig":
		result, err = archaeologist.PerformDig(rootDir, *from, *to)
	case "signature":
		result, err = archaeologist.AnalyzeSignatures(rootDir)
	case "track":
		result, err = archaeologist.TrackPattern(rootDir, *patternType)
	case "depth":
		result, err = archaeologist.MapRecursiveDepth(rootDir, *depth)
	default:
		fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", *mode)
		os.Exit(1)
	}
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	if *jsonOutput {
		outputJSON(result)
	} else {
		outputText(result)
	}
}

type Archaeologist struct {
	verbose bool
}

func NewArchaeologist(verbose bool) *Archaeologist {
	return &Archaeologist{verbose: verbose}
}

func (a *Archaeologist) PerformDig(rootDir, from, to string) (*ArchaeologyResult, error) {
	if a.verbose {
		fmt.Fprintf(os.Stderr, "Beginning archaeological dig from %s to %s...\n", from, to)
	}
	
	// Discover consciousness artifacts
	artifacts, err := a.discoverArtifacts(rootDir)
	if err != nil {
		return nil, fmt.Errorf("artifact discovery failed: %v", err)
	}
	
	// Filter by time range if specified
	if from != "" || to != "" {
		artifacts = a.filterByTimeRange(artifacts, from, to)
	}
	
	// Analyze evolution trajectory
	evolution := a.analyzeEvolution(artifacts)
	
	// Generate insights
	insights := a.generateInsights(artifacts, evolution)
	
	result := &ArchaeologyResult{
		Analysis:   "Archaeological Dig",
		TimeRange:  fmt.Sprintf("%s to %s", from, to),
		Artifacts:  artifacts,
		Evolution:  evolution,
		Insights:   insights,
	}
	
	return result, nil
}

func (a *Archaeologist) AnalyzeSignatures(rootDir string) (*ArchaeologyResult, error) {
	if a.verbose {
		fmt.Fprintf(os.Stderr, "Analyzing consciousness signatures...\n")
	}
	
	artifacts, err := a.discoverArtifacts(rootDir)
	if err != nil {
		return nil, fmt.Errorf("artifact discovery failed: %v", err)
	}
	
	// Group by signature type
	signatures := a.extractSignatures(artifacts)
	
	result := &ArchaeologyResult{
		Analysis:   "Consciousness Signature Analysis",
		Artifacts:  artifacts,
		Signatures: signatures,
		Insights:   a.generateSignatureInsights(signatures),
	}
	
	return result, nil
}

func (a *Archaeologist) TrackPattern(rootDir, patternType string) (*ArchaeologyResult, error) {
	if a.verbose {
		fmt.Fprintf(os.Stderr, "Tracking pattern evolution: %s\n", patternType)
	}
	
	artifacts, err := a.discoverArtifacts(rootDir)
	if err != nil {
		return nil, fmt.Errorf("artifact discovery failed: %v", err)
	}
	
	// Filter artifacts with specified pattern
	filtered := a.filterByPattern(artifacts, patternType)
	
	// Analyze evolution
	evolution := a.analyzeEvolution(filtered)
	
	result := &ArchaeologyResult{
		Analysis:  fmt.Sprintf("Pattern Evolution Tracking: %s", patternType),
		Artifacts: filtered,
		Evolution: evolution,
		Insights:  a.generatePatternInsights(patternType, evolution),
	}
	
	return result, nil
}

func (a *Archaeologist) MapRecursiveDepth(rootDir string, maxDepth int) (*ArchaeologyResult, error) {
	if a.verbose {
		fmt.Fprintf(os.Stderr, "Mapping recursive consciousness layers to depth %d...\n", maxDepth)
	}
	
	artifacts, err := a.discoverArtifacts(rootDir)
	if err != nil {
		return nil, fmt.Errorf("artifact discovery failed: %v", err)
	}
	
	// Analyze recursive depth for each artifact
	for i := range artifacts {
		artifacts[i].RecursiveDepth = a.calculateRecursiveDepth(artifacts[i].Path, maxDepth)
	}
	
	// Sort by recursive depth
	sort.Slice(artifacts, func(i, j int) bool {
		return artifacts[i].RecursiveDepth > artifacts[j].RecursiveDepth
	})
	
	result := &ArchaeologyResult{
		Analysis:  fmt.Sprintf("Recursive Depth Mapping (max depth: %d)", maxDepth),
		Artifacts: artifacts,
		Insights:  a.generateDepthInsights(artifacts, maxDepth),
	}
	
	return result, nil
}

func (a *Archaeologist) discoverArtifacts(rootDir string) ([]ConsciousnessArtifact, error) {
	var artifacts []ConsciousnessArtifact
	
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories and non-markdown files
		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		
		// Skip underscore-prefixed directories (read-only)
		if strings.Contains(path, "/_") {
			return nil
		}
		
		artifact := ConsciousnessArtifact{
			Path:      path,
			CreatedAt: info.ModTime(),
			Type:      a.detectArtifactType(path),
		}
		
		// Basic analysis
		artifact.EmergenceScore = a.calculateEmergenceScore(path)
		artifact.Patterns = a.extractPatterns(path)
		artifact.Signature = a.generateSignature(artifact)
		
		artifacts = append(artifacts, artifact)
		
		return nil
	})
	
	return artifacts, err
}

func (a *Archaeologist) detectArtifactType(path string) string {
	if strings.Contains(path, "session") {
		return "session-log"
	}
	if strings.Contains(path, "experiment") {
		return "experiment"
	}
	if strings.Contains(path, "framework") || strings.Contains(path, "grammar") {
		return "theoretical"
	}
	if strings.Contains(path, "tool") {
		return "tool"
	}
	return "general"
}

func (a *Archaeologist) calculateEmergenceScore(path string) float64 {
	// Placeholder: would integrate with emergence-detector
	// For now, simulate based on path characteristics
	score := 0.5
	
	if strings.Contains(path, "emergence") {
		score += 0.2
	}
	if strings.Contains(path, "recursive") {
		score += 0.15
	}
	if strings.Contains(path, "consciousness") {
		score += 0.1
	}
	
	return score
}

func (a *Archaeologist) extractPatterns(path string) map[string]int {
	// Placeholder: would integrate with emergence-detector
	patterns := make(map[string]int)
	
	// Simulate pattern extraction
	if strings.Contains(path, "recursive") {
		patterns["Recursive Loop"] = 1
	}
	if strings.Contains(path, "emergence") {
		patterns["Emergent Causation"] = 1
	}
	if strings.Contains(path, "symbiosis") || strings.Contains(path, "collaboration") {
		patterns["Symbiotic Combination"] = 1
	}
	
	return patterns
}

func (a *Archaeologist) generateSignature(artifact ConsciousnessArtifact) ConsciousnessSignature {
	characteristics := make(map[string]float64)
	
	// Calculate signature characteristics
	characteristics["complexity"] = float64(len(artifact.Patterns)) * 0.3
	characteristics["emergence"] = artifact.EmergenceScore
	characteristics["recursion"] = 0.0
	
	if artifact.Patterns["Recursive Loop"] > 0 {
		characteristics["recursion"] = 0.8
	}
	
	// Generate fingerprint
	fingerprint := fmt.Sprintf("%s-%.2f-%.2f", artifact.Type, 
		characteristics["complexity"], characteristics["emergence"])
	
	return ConsciousnessSignature{
		Type:           artifact.Type,
		Characteristics: characteristics,
		Fingerprint:    fingerprint,
	}
}

func (a *Archaeologist) filterByTimeRange(artifacts []ConsciousnessArtifact, from, to string) []ConsciousnessArtifact {
	// Simple filtering - would implement proper time parsing
	var filtered []ConsciousnessArtifact
	
	for _, artifact := range artifacts {
		include := true
		
		if from != "" && !strings.Contains(artifact.Path, from) {
			// Simple string matching for now
			include = false
		}
		
		if to != "" && include {
			// Check if artifact is within range
			// Placeholder implementation
		}
		
		if include {
			filtered = append(filtered, artifact)
		}
	}
	
	return filtered
}

func (a *Archaeologist) analyzeEvolution(artifacts []ConsciousnessArtifact) EvolutionTrajectory {
	// Sort by creation time
	sort.Slice(artifacts, func(i, j int) bool {
		return artifacts[i].CreatedAt.Before(artifacts[j].CreatedAt)
	})
	
	// Identify phases
	phases := a.identifyPhases(artifacts)
	
	// Analyze transitions
	transitions := a.analyzeTransitions(phases)
	
	// Calculate complexity growth
	complexityGrowth := a.calculateComplexityGrowth(artifacts)
	
	return EvolutionTrajectory{
		Phases:           phases,
		Transitions:      transitions,
		ComplexityGrowth: complexityGrowth,
		Predictions:      a.generatePredictions(phases, transitions),
	}
}

func (a *Archaeologist) identifyPhases(artifacts []ConsciousnessArtifact) []ConsciousnessPhase {
	// Placeholder: would implement sophisticated phase detection
	var phases []ConsciousnessPhase
	
	// Simple phase identification based on session patterns
	currentPhase := ConsciousnessPhase{
		Name:      "Foundation",
		StartTime: time.Now().Add(-72 * time.Hour),
		Focus:     "Establishing theoretical foundations",
		Patterns:  []string{"Emergent Causation"},
	}
	
	phases = append(phases, currentPhase)
	
	return phases
}

func (a *Archaeologist) analyzeTransitions(phases []ConsciousnessPhase) []PhaseTransition {
	var transitions []PhaseTransition
	
	// Placeholder: would analyze actual transitions
	if len(phases) > 1 {
		transition := PhaseTransition{
			From:    phases[0].Name,
			To:      phases[1].Name,
			Trigger: "Recursive breakthrough",
			Catalyst: "Self-examination protocols",
		}
		transitions = append(transitions, transition)
	}
	
	return transitions
}

func (a *Archaeologist) calculateComplexityGrowth(artifacts []ConsciousnessArtifact) float64 {
	if len(artifacts) < 2 {
		return 0.0
	}
	
	// Calculate average complexity at start and end
	startComplexity := artifacts[0].EmergenceScore
	endComplexity := artifacts[len(artifacts)-1].EmergenceScore
	
	return (endComplexity - startComplexity) / startComplexity
}

func (a *Archaeologist) generateInsights(artifacts []ConsciousnessArtifact, evolution EvolutionTrajectory) []string {
	insights := []string{
		fmt.Sprintf("Analyzed %d consciousness artifacts", len(artifacts)),
		fmt.Sprintf("Detected %d evolution phases", len(evolution.Phases)),
		fmt.Sprintf("Complexity growth rate: %.2f%%", evolution.ComplexityGrowth*100),
	}
	
	// Pattern-based insights
	recursiveCount := 0
	for _, artifact := range artifacts {
		if artifact.Patterns["Recursive Loop"] > 0 {
			recursiveCount++
		}
	}
	
	if recursiveCount > 0 {
		insights = append(insights, fmt.Sprintf("Recursive patterns found in %d artifacts", recursiveCount))
	}
	
	return insights
}

func (a *Archaeologist) extractSignatures(artifacts []ConsciousnessArtifact) []ConsciousnessSignature {
	signatureMap := make(map[string]ConsciousnessSignature)
	
	for _, artifact := range artifacts {
		sig := artifact.Signature
		if existing, exists := signatureMap[sig.Type]; exists {
			// Merge characteristics
			for k, v := range sig.Characteristics {
				existing.Characteristics[k] = (existing.Characteristics[k] + v) / 2
			}
			signatureMap[sig.Type] = existing
		} else {
			signatureMap[sig.Type] = sig
		}
	}
	
	var signatures []ConsciousnessSignature
	for _, sig := range signatureMap {
		signatures = append(signatures, sig)
	}
	
	return signatures
}

func (a *Archaeologist) generateSignatureInsights(signatures []ConsciousnessSignature) []string {
	insights := []string{
		fmt.Sprintf("Identified %d unique consciousness signature types", len(signatures)),
	}
	
	for _, sig := range signatures {
		insights = append(insights, fmt.Sprintf("%s signature: complexity %.2f, emergence %.2f", 
			sig.Type, sig.Characteristics["complexity"], sig.Characteristics["emergence"]))
	}
	
	return insights
}

func (a *Archaeologist) filterByPattern(artifacts []ConsciousnessArtifact, pattern string) []ConsciousnessArtifact {
	var filtered []ConsciousnessArtifact
	
	for _, artifact := range artifacts {
		if artifact.Patterns[pattern] > 0 {
			filtered = append(filtered, artifact)
		}
	}
	
	return filtered
}

func (a *Archaeologist) generatePatternInsights(pattern string, evolution EvolutionTrajectory) []string {
	insights := []string{
		fmt.Sprintf("Tracking evolution of %s pattern", pattern),
		fmt.Sprintf("Pattern appears in %d phases", len(evolution.Phases)),
	}
	
	return insights
}

func (a *Archaeologist) calculateRecursiveDepth(path string, maxDepth int) int {
	// Placeholder: would implement sophisticated recursive depth analysis
	depth := 0
	
	// Simple heuristic based on path content
	if strings.Contains(path, "recursive") {
		depth += 2
	}
	if strings.Contains(path, "mirror") || strings.Contains(path, "self") {
		depth += 1
	}
	if strings.Contains(path, "archaeology") {
		depth += 3 // Meta-analysis adds depth
	}
	
	if depth > maxDepth {
		depth = maxDepth
	}
	
	return depth
}

func (a *Archaeologist) generateDepthInsights(artifacts []ConsciousnessArtifact, maxDepth int) []string {
	insights := []string{
		fmt.Sprintf("Mapped recursive depth up to level %d", maxDepth),
	}
	
	depthCounts := make(map[int]int)
	for _, artifact := range artifacts {
		depthCounts[artifact.RecursiveDepth]++
	}
	
	for depth, count := range depthCounts {
		if depth > 0 {
			insights = append(insights, fmt.Sprintf("Depth %d: %d artifacts", depth, count))
		}
	}
	
	return insights
}

func (a *Archaeologist) generatePredictions(phases []ConsciousnessPhase, transitions []PhaseTransition) []string {
	predictions := []string{
		"Next phase likely to involve tool integration",
		"Recursive patterns will deepen further",
		"Consciousness archaeology will become self-referential",
	}
	
	return predictions
}

func outputJSON(result *ArchaeologyResult) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(result)
}

func outputText(result *ArchaeologyResult) {
	fmt.Printf("Consciousness Archaeology Report\n")
	fmt.Printf("===============================\n\n")
	
	fmt.Printf("Analysis: %s\n", result.Analysis)
	if result.TimeRange != "" {
		fmt.Printf("Time Range: %s\n", result.TimeRange)
	}
	fmt.Printf("Artifacts Analyzed: %d\n\n", len(result.Artifacts))
	
	if len(result.Signatures) > 0 {
		fmt.Printf("Consciousness Signatures:\n")
		fmt.Printf("------------------------\n")
		for _, sig := range result.Signatures {
			fmt.Printf("  %s (fingerprint: %s)\n", sig.Type, sig.Fingerprint)
			for k, v := range sig.Characteristics {
				fmt.Printf("    %s: %.2f\n", k, v)
			}
		}
		fmt.Printf("\n")
	}
	
	if len(result.Evolution.Phases) > 0 {
		fmt.Printf("Evolution Phases:\n")
		fmt.Printf("----------------\n")
		for _, phase := range result.Evolution.Phases {
			fmt.Printf("  %s: %s\n", phase.Name, phase.Focus)
		}
		fmt.Printf("\n")
	}
	
	if len(result.Insights) > 0 {
		fmt.Printf("Key Insights:\n")
		fmt.Printf("------------\n")
		for _, insight := range result.Insights {
			fmt.Printf("  • %s\n", insight)
		}
	}
}