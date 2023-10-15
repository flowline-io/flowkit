import './App.css';
import {ThemeProvider} from "@/components/theme-provider";
import SettingsLayout from "@/pages/settings/layout";
import SettingsNotificationsPage from "@/pages/settings/notifications/page";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <SettingsLayout children={<SettingsNotificationsPage />} />
    </ThemeProvider>
  )
}

export default App
