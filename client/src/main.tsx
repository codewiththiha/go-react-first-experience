// src/main.tsx
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import { Provider } from "@/components/ui/provider";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// Use relative URL in production, absolute in development
export const BASEURL = import.meta.env.DEV
	? "http://127.0.0.1:4000/api"
	: "/api";
// force production mode

const queryClient = new QueryClient();
createRoot(document.getElementById("root")!).render(
	<StrictMode>
		<QueryClientProvider client={queryClient}>
			<Provider>
				<App />
			</Provider>
		</QueryClientProvider>
	</StrictMode>
);
