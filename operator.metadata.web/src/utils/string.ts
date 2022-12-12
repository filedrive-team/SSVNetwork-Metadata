export const ellipsisByLength = (str = '', length = 6) => {
  return str.length > 2 * length
    ? `${str.slice(0, length)}***${str.slice(-length)}`
    : str;
};
