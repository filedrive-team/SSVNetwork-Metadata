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
    name: 'detail',
    path: RouterPath.detail,
    auth: true,
    component: lazy(() => import('@/pages/Detail')),
  },
  {
    name: 'profile',
    path: RouterPath.profile,
    auth: false,
    component: lazy(() => import('@/pages/Profile')),
  },
  {
    name: 'auth',
    path: RouterPath.auth,
    auth: false,
    component: lazy(() => import('@/pages/Auth')),
  },
  {
    name: '404',
    path: RouterPath.error,
    auth: false,
    component: lazy(() => import('@/404')),
  },
];

export { RouterPath, Routes };
