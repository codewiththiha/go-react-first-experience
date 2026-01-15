// src/components/ui/provider.tsx
import { ChakraProvider, defaultSystem } from "@chakra-ui/react";
import { ColorModeProvider } from "@/components/ui/color-mode";

export function Provider({ children }: { children: React.ReactNode }) {
	return (
		<ChakraProvider value={defaultSystem /* or your `system` */}>
			<ColorModeProvider>{children}</ColorModeProvider>
		</ChakraProvider>
	);
}
