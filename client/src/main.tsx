// src/main.tsx
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import { Provider } from "@/components/ui/provider";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// uncomment this if you are in development stage
// export const BASEURL =
// 	import.meta.env.MODE === "development" ? "http://127.0.0.1:4000/api" : "/api";

// fix 2 isn't work either
// export const BASEURL = import.meta.env.PROD
// 	? "/api"
// 	: "http://127.0.0.1:4000/api";

// force production mode
export const BASEURL = "/api";

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
