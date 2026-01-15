// src/components/navbar.tsx
import { Box, Container, Flex, Heading, IconButton } from "@chakra-ui/react";
import { useColorMode } from "@/components/ui/color-mode";
import { Tooltip } from "@/components/ui/tooltip"; // Use the CLI-generated snippet
import { LuSun, LuMoon } from "react-icons/lu"; // Recommended for cleaner code

export default function Navbar() {
	const { colorMode, toggleColorMode } = useColorMode();
	const isDark = colorMode === "dark";

	return (
		<Box
			as="nav"
			position="sticky"
			top={0}
			zIndex={10}
			bg="bg"
			borderBottomWidth="1px"
			borderColor="border"
		>
			<Container maxW="container.lg" py={2}>
				<Flex align="center" justify="space-between" gap={4}>
					<Heading as="h1" size="md">
						TODO
					</Heading>

					<Tooltip
						content={isDark ? "Switch to light mode" : "Switch to dark mode"}
					>
						<IconButton
							aria-label="Toggle color mode"
							variant="ghost"
							size="sm"
							onClick={toggleColorMode}
						>
							{isDark ? <LuSun /> : <LuMoon />}
						</IconButton>
					</Tooltip>
				</Flex>
			</Container>
		</Box>
	);
}
