import './App.css';
import {ThemeProvider} from "@/components/theme-provider";
import DashboardPage from "@/pages/dashboard/Dashboard";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <DashboardPage />
    </ThemeProvider>
  )
}

export default App
