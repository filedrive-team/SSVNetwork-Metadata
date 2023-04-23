import { lazy } from 'react';
import { RouteModel } from '@/models/RouteModel';

enum RouterPath {
  home = '/',
  detail = '/detail',
  profile = '/profile',
  error = '/error',
  auth = '/auth',
}

const Routes: RouteModel[] = [
  {
    name: 'Home',
    path: RouterPath.home,
    auth: false,
    component: lazy(() => import('@/pages/Home')),
  },
  {
    name: 'profile',
    path: RouterPath.profile,
    auth: false,
    component: lazy(() => import('@/pages/Profile')),
  },
  {
    name: '404',
    path: RouterPath.error,
    auth: false,
    component: lazy(() => import('@/404')),
  },
];

export { RouterPath, Routes };
