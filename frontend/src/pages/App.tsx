import './App.css';
import {ThemeProvider} from "@/components/theme-provider";
import SettingsLayout from "@/pages/settings/layout";
import SettingsNotificationsPage from "@/pages/settings/notifications/page";
import DashboardPage from "@/pages/dashboard/Dashboard";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      {/*<SettingsLayout children={<SettingsNotificationsPage />} />*/}
      <DashboardPage />
    </ThemeProvider>
  )
}

export default App
