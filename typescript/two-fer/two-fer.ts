export function twoFer(name?: string ): string {
  const defaultMessage = 'One for you, one for me.'
  if (typeof(name) === 'undefined') {
    return defaultMessage
  }
  return `One for ${name}, one for me.`
}
