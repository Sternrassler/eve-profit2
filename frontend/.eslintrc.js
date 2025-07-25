module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    '@typescript-eslint/recommended',
    'prettier'
  ],
  ignorePatterns: ['dist', '.eslintrc.js'],
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint', 'prettier'],
  rules: {
    // Universal Development Guidelines - Clean Code Standards
    'prettier/prettier': 'error',
    '@typescript-eslint/no-unused-vars': 'error',
    '@typescript-eslint/explicit-function-return-type': 'warn',
    '@typescript-eslint/no-explicit-any': 'warn',
    
    // EVE Project Specific Rules
    'prefer-const': 'error',
    'no-var': 'error',
    'camelcase': 'error', // SonarQube compliance
    
    // Clean Code: Meaningful Names
    'id-length': ['error', { min: 2, exceptions: ['i', 'j', 'x', 'y'] }],
    
    // Clean Code: Functions should be small
    'max-lines-per-function': ['warn', { max: 20 }],
    'max-params': ['warn', { max: 4 }]
  },
};
