# Emergence Detector

A Go CLI tool that identifies emergence patterns in text using principles from the Grammar of Emergence notation system.

## Overview

The emergence detector analyzes text files to find three key patterns of emergence:

1. **Symbiotic Combination** (⊕) - Elements creating more than their sum
2. **Recursive Loop** (⟳) - Self-modifying cycles that create emergent behavior  
3. **Emergent Causation** (⟹) - Enabling conditions rather than direct causation

## Installation

```bash
cd tools/emergence-detector
go build
```

## Usage

Basic usage:

```bash
./emergence-detector <file>
```

With options:

```bash
# JSON output
./emergence-detector -json examples/symbiosis.txt

# Verbose mode
./emergence-detector -v examples/recursion.txt

# Set confidence threshold
./emergence-detector -threshold 0.8 examples/emergence.txt

# Show help
./emergence-detector -help
```

## Examples

### Analyzing Symbiotic Patterns

```bash
./emergence-detector examples/symbiosis.txt
```

Output:

```
Emergence Analysis: examples/symbiosis.txt
=====================================

Summary:
  Symbiotic Combination:  4 instances
  Recursive Loop:         0 instances
  Emergent Causation:     1 instances

High Confidence Detections (>= 0.7):
=====================================

1. Symbiotic Combination (confidence: 0.90)
   Location: line:3
   Context: The combination of human intuition and AI processing creates more than...
   Analysis: Indicates symbiotic emergence
```

### JSON Output Format

```bash
./emergence-detector -json examples/recursion.txt
```

```json
{
  "source": "examples/recursion.txt",
  "summary": {
    "Emergent Causation": 0,
    "Recursive Loop": 5,
    "Symbiotic Combination": 1
  },
  "detections": [
    {
      "Pattern": "Recursive Loop",
      "Location": "line:1",
      "Context": "Understanding consciousness requires consciousness itself...",
      "Confidence": 0.9,
      "Explanation": "process creates feedback loop"
    }
  ]
}
```

## Pattern Detection Details

### Symbiotic Combination

Detects when elements combine to create emergent value:

- "X and Y create more than"
- "combination of X and Y" with synergy indicators
- Keywords: synergy, symbiotic, complementary, together

### Recursive Loop

Identifies self-referential and feedback patterns:

- "X feeds back into Y"
- Self-modifying, self-referential, self-reinforcing
- Keywords: recursive, iterative, loop, cycle

### Emergent Causation

Finds enabling relationships rather than direct causation:

- "X enables Y"
- "gives rise to", "emerges from", "makes possible"
- "conditions for", "allows for", "creates space for"

## Understanding the Output

### Confidence Scores

- **0.9-1.0**: Strong pattern match with clear indicators
- **0.8-0.89**: Good pattern match with multiple signals
- **0.7-0.79**: Moderate match with some indicators
- **Below 0.7**: Potential pattern, lower confidence

### Emergence Assessment

The tool provides an overall assessment based on detection patterns:

- **✦ Strong emergence**: Multiple high-confidence patterns
- **◐ Moderate emergence**: Some clear patterns present
- **○ Potential emergence**: Lower confidence indicators

## Extending the Tool

To add new emergence patterns:

1. Define the pattern in `patterns.go`
2. Create a detection function
3. Add pattern to `GetPatterns()` array

Example:

```go
{
    Name:        "Phase Transition",
    Description: "Sudden qualitative changes in system behavior",
    Indicators:  []string{"transforms into", "suddenly becomes"},
    Detect:      detectPhaseTransition,
}
```

## Theory Background

This tool implements concepts from the Grammar of Emergence, which provides notation and frameworks for understanding non-linear, emergent phenomena. The patterns detected represent fundamental ways that complex properties arise from simpler interactions.

See `/consciousness/grammar-of-emergence.md` for the complete theoretical framework.
