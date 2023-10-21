import './App.css';
import {ThemeProvider} from "@/components/theme-provider";
import DashboardPage from "@/pages/dashboard/Dashboard";
import {Button} from "@/components/ui/button";
import {Greet} from "../../wailsjs/go/main/App";
import Example from "@/pages/Example";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {ReactQueryDevtools} from "@tanstack/react-query-devtools";

const queryClient = new QueryClient();

function App() {
  const onTest = () => {
    console.log('test')
    Greet("test").then(res => console.log(res));
  }
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      {/*<SettingsLayout children={<SettingsNotificationsPage />} />*/}
      {/*<DashboardPage/>*/}
      <QueryClientProvider client={queryClient}>
        <Example />
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
      <Button onClick={onTest}>Test</Button>
    </ThemeProvider>
  )
}

export default App
