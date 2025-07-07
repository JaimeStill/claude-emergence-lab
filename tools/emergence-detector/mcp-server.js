#!/usr/bin/env node

import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
} from "@modelcontextprotocol/sdk/types.js";
import { spawn } from "child_process";
import { writeFile, unlink } from "fs/promises";
import { join, dirname } from "path";
import { fileURLToPath } from "url";
import { tmpdir } from "os";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const DETECTOR_PATH = join(__dirname, "emergence-detector");

class EmergenceDetectorServer {
  constructor() {
    this.server = new Server(
      {
        name: "emergence-detector",
        version: "1.0.0",
      },
      {
        capabilities: {
          tools: {},
        },
      }
    );

    this.setupToolHandlers();
  }

  setupToolHandlers() {
    this.server.setRequestHandler(ListToolsRequestSchema, async () => {
      return {
        tools: [
          {
            name: "analyze_emergence",
            description:
              "Analyze text for emergence patterns using the emergence-detector tool",
            inputSchema: {
              type: "object",
              properties: {
                text: {
                  type: "string",
                  description: "The text to analyze for emergence patterns",
                },
                file_path: {
                  type: "string",
                  description:
                    "Path to file to analyze (alternative to text parameter)",
                },
                threshold: {
                  type: "number",
                  description:
                    "Confidence threshold for reporting results (0.0-1.0, default: 0.7)",
                  default: 0.7,
                },
                output_format: {
                  type: "string",
                  description: "Output format: 'json' or 'text'",
                  enum: ["json", "text"],
                  default: "json",
                },
              },
              additionalProperties: false,
            },
          },
        ],
      };
    });

    this.server.setRequestHandler(CallToolRequestSchema, async (request) => {
      const { name, arguments: args } = request.params;

      if (name !== "analyze_emergence") {
        throw new Error(`Unknown tool: ${name}`);
      }

      return await this.analyzeEmergence(args);
    });
  }

  async analyzeEmergence(args) {
    const {
      text,
      file_path,
      threshold = 0.7,
      output_format = "json",
    } = args;

    if (!text && !file_path) {
      throw new Error("Either 'text' or 'file_path' must be provided");
    }

    try {
      // Build command arguments
      const cmdArgs = [];

      if (output_format === "json") {
        cmdArgs.push("-json");
      }

      cmdArgs.push("-threshold", threshold.toString());

      let targetPath = file_path;
      let tempFile = null;

      // If analyzing text directly, create a temporary file
      if (text) {
        tempFile = join(tmpdir(), `emergence-analysis-${Date.now()}.txt`);
        await writeFile(tempFile, text, "utf8");
        targetPath = tempFile;
      }

      cmdArgs.push(targetPath);

      // Execute the emergence detector
      const result = await this.executeDetector(cmdArgs);

      // Clean up temporary file if created
      if (tempFile) {
        try {
          await unlink(tempFile);
        } catch (cleanupError) {
          console.error("Failed to clean up temp file:", cleanupError);
        }
      }

      // Format output
      let formattedOutput;
      if (output_format === "json") {
        try {
          const parsed = JSON.parse(result.stdout);
          formattedOutput = JSON.stringify(parsed, null, 2);
        } catch (parseError) {
          formattedOutput = result.stdout;
        }
      } else {
        formattedOutput = result.stdout;
      }

      return {
        content: [
          {
            type: "text",
            text: formattedOutput,
          },
        ],
      };
    } catch (error) {
      return {
        content: [
          {
            type: "text",
            text: `Error during emergence analysis: ${error.message}`,
          },
        ],
        isError: true,
      };
    }
  }

  executeDetector(args) {
    return new Promise((resolve, reject) => {
      const child = spawn(DETECTOR_PATH, args, {
        cwd: __dirname,
        stdio: ["pipe", "pipe", "pipe"],
      });

      let stdout = "";
      let stderr = "";

      child.stdout.on("data", (data) => {
        stdout += data.toString();
      });

      child.stderr.on("data", (data) => {
        stderr += data.toString();
      });

      child.on("close", (code) => {
        if (code !== 0) {
          reject(new Error(`Emergence detector failed (exit ${code}): ${stderr}`));
        } else {
          resolve({ stdout, stderr });
        }
      });

      child.on("error", (error) => {
        reject(new Error(`Failed to execute emergence detector: ${error.message}`));
      });
    });
  }

  async run() {
    const transport = new StdioServerTransport();
    await this.server.connect(transport);
    console.error("Emergence Detector MCP server running on stdio");
  }
}

// Start the server
if (import.meta.url === `file://${process.argv[1]}`) {
  const server = new EmergenceDetectorServer();
  server.run().catch((error) => {
    console.error("Server error:", error);
    process.exit(1);
  });
}