import './App.css';
import {ThemeProvider} from "@/components/theme-provider";
import DashboardPage from "@/pages/dashboard/Dashboard";
import {Button} from "@/components/ui/button";
import {Greet} from "../../wailsjs/go/main/App";

function App() {
  const onTest = () => {
    console.log('test')
    Greet("test").then(res => console.log(res));
  }
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      {/*<SettingsLayout children={<SettingsNotificationsPage />} />*/}
      <DashboardPage/>
      <Button onClick={onTest}>Test</Button>
    </ThemeProvider>
  )
}

export default App
