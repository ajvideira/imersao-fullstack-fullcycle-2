import { AppBar, IconButton, Toolbar, Typography } from '@material-ui/core';
import DriverIcon from '@material-ui/icons/DriveEta';

const Sidebar: React.FC = () => {
  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu">
          <DriverIcon />
        </IconButton>
        <Typography variant="h6">Code Delivery</Typography>
      </Toolbar>
    </AppBar>
  );
};
export default Sidebar;
