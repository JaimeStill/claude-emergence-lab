package main

import (
	"fmt"
	"regexp"
	"strings"
)

// EmergencePattern represents a type of emergence we can detect
type EmergencePattern struct {
	Name        string
	Description string
	Indicators  []string
	Detect      func(text string) []Detection
}

// Detection represents a found instance of emergence
type Detection struct {
	Pattern     string
	Location    string
	Context     string
	Confidence  float64
	Explanation string
}

// GetPatterns returns all emergence patterns we can detect
func GetPatterns() []EmergencePattern {
	return []EmergencePattern{
		{
			Name:        "Symbiotic Combination",
			Description: "Elements combining to create more than their sum (A ⊕ B)",
			Indicators:  []string{"more than", "greater than the sum", "synergy", "combined effect", "together create"},
			Detect:      detectSymbiosis,
		},
		{
			Name:        "Recursive Loop",
			Description: "Self-modifying cycles that create emergent behavior (⟳)",
			Indicators:  []string{"feeds back", "recursive", "self-modifying", "iterative", "reinforcing"},
			Detect:      detectRecursion,
		},
		{
			Name:        "Emergent Causation",
			Description: "Enabling conditions rather than direct causation (⟹)",
			Indicators:  []string{"enables", "gives rise to", "emerges from", "conditions for", "makes possible"},
			Detect:      detectEmergentCausation,
		},
	}
}

// detectSymbiosis finds instances of symbiotic combination
func detectSymbiosis(text string) []Detection {
	var detections []Detection
	lines := strings.Split(text, "\n")
	
	// Pattern 1: "X and Y create/produce/yield more than"
	pattern1 := regexp.MustCompile(`(\w+)\s+and\s+(\w+)\s+(create|produce|yield|generate)\s+more than`)
	
	// Pattern 2: "combination of X and Y"
	pattern2 := regexp.MustCompile(`combination of\s+(\w+)\s+and\s+(\w+)`)
	
	// Pattern 3: synergy keywords
	synergyPattern := regexp.MustCompile(`(synerg|symbiotic|complementary|together)`)
	
	for i, line := range lines {
		if match := pattern1.FindStringSubmatch(line); match != nil {
			detections = append(detections, Detection{
				Pattern:     "Symbiotic Combination",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.9,
				Explanation: formatExplanation(match[1], match[2], "create symbiotic value"),
			})
		}
		
		if match := pattern2.FindStringSubmatch(line); match != nil && synergyPattern.MatchString(line) {
			detections = append(detections, Detection{
				Pattern:     "Symbiotic Combination",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.8,
				Explanation: formatExplanation(match[1], match[2], "combine symbiotically"),
			})
		}
		
		// Check for indicator phrases
		for _, indicator := range []string{"more than", "greater than the sum"} {
			if strings.Contains(strings.ToLower(line), indicator) {
				detections = append(detections, Detection{
					Pattern:     "Symbiotic Combination",
					Location:    formatLocation(i+1),
					Context:     line,
					Confidence:  0.7,
					Explanation: "Indicates symbiotic emergence",
				})
				break
			}
		}
	}
	
	return detections
}

// detectRecursion finds recursive patterns
func detectRecursion(text string) []Detection {
	var detections []Detection
	lines := strings.Split(text, "\n")
	
	// Pattern 1: "X feeds back into Y"
	feedbackPattern := regexp.MustCompile(`(\w+)\s+feeds?\s+back\s+into`)
	
	// Pattern 2: self-reference
	selfPattern := regexp.MustCompile(`self[- ]?(modifying|referential|reinforcing|amplifying)`)
	
	// Pattern 3: iterative/recursive keywords
	recursivePattern := regexp.MustCompile(`(recursive|iterative|loop|cycle)`)
	
	for i, line := range lines {
		if match := feedbackPattern.FindStringSubmatch(line); match != nil {
			detections = append(detections, Detection{
				Pattern:     "Recursive Loop",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.9,
				Explanation: match[1] + " creates feedback loop",
			})
		}
		
		if selfPattern.MatchString(line) {
			detections = append(detections, Detection{
				Pattern:     "Recursive Loop",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.85,
				Explanation: "Self-modifying behavior detected",
			})
		}
		
		if recursivePattern.MatchString(strings.ToLower(line)) {
			detections = append(detections, Detection{
				Pattern:     "Recursive Loop",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.7,
				Explanation: "Recursive pattern indicated",
			})
		}
	}
	
	return detections
}

// detectEmergentCausation finds emergent causation patterns
func detectEmergentCausation(text string) []Detection {
	var detections []Detection
	lines := strings.Split(text, "\n")
	
	// Pattern 1: "X enables Y"
	enablePattern := regexp.MustCompile(`(\w+)\s+enables?\s+(\w+)`)
	
	// Pattern 2: "gives rise to" or "emerges from"
	emergencePattern := regexp.MustCompile(`(gives?\s+rise\s+to|emerges?\s+from|makes?\s+possible)`)
	
	// Pattern 3: condition-based language
	conditionPattern := regexp.MustCompile(`(conditions?\s+for|allows?\s+for|creates?\s+space\s+for)`)
	
	for i, line := range lines {
		if match := enablePattern.FindStringSubmatch(line); match != nil {
			detections = append(detections, Detection{
				Pattern:     "Emergent Causation",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.85,
				Explanation: match[1] + " enables emergence of " + match[2],
			})
		}
		
		if emergencePattern.MatchString(line) {
			detections = append(detections, Detection{
				Pattern:     "Emergent Causation",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.9,
				Explanation: "Emergent relationship identified",
			})
		}
		
		if conditionPattern.MatchString(line) {
			detections = append(detections, Detection{
				Pattern:     "Emergent Causation",
				Location:    formatLocation(i+1),
				Context:     line,
				Confidence:  0.75,
				Explanation: "Enabling conditions present",
			})
		}
	}
	
	return detections
}

// Helper functions
func formatLocation(line int) string {
	return fmt.Sprintf("line:%d", line)
}

func formatExplanation(elem1, elem2, action string) string {
	return elem1 + " and " + elem2 + " " + action
}