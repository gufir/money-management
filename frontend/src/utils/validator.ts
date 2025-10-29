export function validateString(value: string, minLength: number, maxLength: number): string | null {
  const n = value.length
  if (n < minLength || n > maxLength) {
    return `Must contain from ${minLength}-${maxLength} characters`
  }
  return null
}

export function validateUsername(value: string): string | null {
  const basic = validateString(value, 3, 100)
  if (basic) return basic

  const isValidUsername = /^[a-z0-9_]+$/.test(value)
  if (!isValidUsername) {
    return 'Must contain only lowercase letters, numbers, and underscores'
  }

  return null
}

export function validatePassword(value: string): string | null {
  const basic = validateString(value, 6, 100)
  if (basic) return basic

  const hasLetter = /[A-Za-z]/.test(value)
  const hasNumber = /[0-9]/.test(value)
  if (!hasLetter || !hasNumber) {
    return 'Must contain at least one letter and one number'
  }

  return null
}

export function validateEmail(value: string): string | null {
  const basic = validateString(value, 3, 100)
  if (basic) return basic

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(value)) {
    return 'Must be a valid email address'
  }

  return null
}

export function validateFullName(value: string): string | null {
  const basic = validateString(value, 3, 100)
  if (basic) return basic

  const isValidFullName = /^[a-zA-Z\s]+$/.test(value)
  if (!isValidFullName) {
    return 'Must contain only letters or spaces'
  }

  return null
}
