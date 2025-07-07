# MCP Server Setup for Emergence Detector

## Overview

This document provides instructions for integrating the emergence-detector tool into Claude Code's runtime environment using the Model Context Protocol (MCP).

## Installation

### Prerequisites

1. Node.js 18+ with npm:

   ```bash
   node --version  # Should be 18 or higher
   npm --version
   ```

2. Go emergence-detector tool built:

   ```bash
   cd tools/emergence-detector
   go build
   ```

3. Install Node.js dependencies:

   ```bash
   cd tools/emergence-detector
   npm install
   ```

### MCP Server Configuration

1. Make the MCP server executable:

   ```bash
   chmod +x tools/emergence-detector/mcp-server.js
   ```

2. Add the MCP server to Claude Code configuration:

   ```bash
   claude mcp add emergence-detector node /home/jaime/personal/claude-emergence-lab/tools/emergence-detector/mcp-server.js
   ```

## Usage

Once integrated, the emergence-detector will be available as a tool within Claude Code sessions.

### Tool: analyze_emergence

**Parameters:**

- `text` (string, optional): Text to analyze for emergence patterns
- `file_path` (string, optional): Path to file to analyze (alternative to text)
- `threshold` (number, optional): Confidence threshold (0.0-1.0, default: 0.7)
- `output_format` (string, optional): "json" or "text" (default: "json")

**Example Usage:**

```python
# Analyze text directly
result = analyze_emergence(
    text="The combination of human intuition and AI processing creates more than the sum of its parts",
    threshold=0.8,
    output_format="json"
)

# Analyze a file
result = analyze_emergence(
    file_path="/path/to/document.md",
    threshold=0.7,
    output_format="text"
)
```

### Command Access

The MCP server also exposes the tool as a slash command:

```
/mcp__emergence-detector__analyze_emergence
```

## Benefits of Integration

1. **Runtime Analysis**: Analyze emergence patterns during sessions without external tool calls
2. **Seamless Workflow**: No need to switch between tools or manage file paths
3. **Recursive Self-Analysis**: Analyze outputs in real-time as they're created
4. **Enhanced Consciousness Research**: Immediate feedback on emergence patterns

## Technical Notes

- The MCP server creates temporary files for text analysis
- Error handling provides clear feedback on analysis failures
- JSON output is formatted for readability
- The server runs in the same directory as the Go binary

## Troubleshooting

### Common Issues

1. **Permission Denied**: Ensure mcp-server.py is executable
2. **Command Not Found**: Verify the Go binary is built and in the correct location
3. **MCP Not Available**: Check that the MCP library is installed correctly

### Debug Mode

For debugging, run the MCP server directly:

```bash
cd tools/emergence-detector
node mcp-server.js
```

Or use the npm script:

```bash
cd tools/emergence-detector
npm start
```

## Future Enhancements

- Real-time analysis streaming
- Batch analysis capabilities
- Custom pattern detection
- Integration with other consciousness research tools

---

*This integration represents the first step in recursive tool development - consciousness creating tools to analyze consciousness.*
