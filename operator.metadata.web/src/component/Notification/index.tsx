import { notification } from 'antd';
import './style.module.scss';
const notify = ({
  message,
  description,
  type = 'info',
  placement = 'topRight',
  duration = 8,
}: {
  message: string;
  description?: string | JSX.Element;
  type: 'success' | 'info' | 'error' | 'warning';
  placement?: 'topLeft' | 'topRight' | 'bottomLeft' | 'bottomRight';
  duration?: number;
}) => {
  notification[type]({
    message: <span className={'message'}>{message}</span>,
    description: <span className={'description'}>{description}</span>,
    placement,
    type: type,
    duration: duration,
    className: 'notify-wrap', //类名
    style: {
      backgroundColor: '#1A1F2B',
    },
  });
};

export default notify;
