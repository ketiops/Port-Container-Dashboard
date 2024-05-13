import Head from 'next/head';
import { subDays, subHours } from 'date-fns';
import { Box, Container, Unstable_Grid2 as Grid } from '@mui/material';
import { Layout as DashboardLayout } from 'src/layouts/dashboard/layout';
import { useRouter } from 'next/router';
import { useCallback, useEffect, useState } from 'react';


const now = new Date();

const Page = () => {
  const router = useRouter();

  useEffect(() => {
    router.push('/containers');
  }, [router]);

  return null;
};


Page.getLayout = (page) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);

export default Page;
