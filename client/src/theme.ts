// src/theme.ts
import { createSystem, defaultConfig } from "@chakra-ui/react";

export const system = createSystem(defaultConfig, {
	theme: {
		tokens: {
			fonts: {
				heading: { value: "Inter, system-ui, sans-serif" },
				body: { value: "Inter, system-ui, sans-serif" },
			},
		},
	},
});
