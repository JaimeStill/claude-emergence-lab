package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// Detector manages the emergence detection process
type Detector struct {
	patterns []EmergencePattern
	verbose  bool
}

// NewDetector creates a new emergence detector
func NewDetector(verbose bool) *Detector {
	return &Detector{
		patterns: GetPatterns(),
		verbose:  verbose,
	}
}

// AnalyzeFile analyzes a file for emergence patterns
func (d *Detector) AnalyzeFile(filename string) (*AnalysisResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	
	return d.AnalyzeText(string(content), filename), nil
}

// AnalyzeText analyzes text for emergence patterns
func (d *Detector) AnalyzeText(text string, source string) *AnalysisResult {
	result := &AnalysisResult{
		Source:     source,
		Detections: make([]Detection, 0),
		Summary:    make(map[string]int),
	}
	
	// Run each pattern detector
	for _, pattern := range d.patterns {
		detections := pattern.Detect(text)
		result.Detections = append(result.Detections, detections...)
		result.Summary[pattern.Name] = len(detections)
		
		if d.verbose && len(detections) > 0 {
			fmt.Printf("Pattern '%s' found %d instances\n", pattern.Name, len(detections))
		}
	}
	
	// Sort detections by confidence
	sort.Slice(result.Detections, func(i, j int) bool {
		return result.Detections[i].Confidence > result.Detections[j].Confidence
	})
	
	return result
}

// AnalysisResult contains the results of emergence detection
type AnalysisResult struct {
	Source     string
	Detections []Detection
	Summary    map[string]int
}

// HasEmergence returns true if any emergence patterns were detected
func (r *AnalysisResult) HasEmergence() bool {
	return len(r.Detections) > 0
}

// GetHighConfidenceDetections returns detections above a confidence threshold
func (r *AnalysisResult) GetHighConfidenceDetections(threshold float64) []Detection {
	var highConf []Detection
	for _, d := range r.Detections {
		if d.Confidence >= threshold {
			highConf = append(highConf, d)
		}
	}
	return highConf
}