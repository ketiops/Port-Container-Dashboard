import { SvgIcon } from '@mui/material';
import { LuContainer } from "react-icons/lu";
import { AiFillSecurityScan } from "react-icons/ai"

export const items = [
  {
    title: 'Container',
    path: '/containers',
    icon: (
      <SvgIcon fontSize="small">
        <LuContainer />
      </SvgIcon>
    )
  },
  {
    title: 'Security Checkpoint',
    path: '/checkpoint',
    icon: (
      <SvgIcon fontSize="small">
        <AiFillSecurityScan />
      </SvgIcon>
    )
  }
];
