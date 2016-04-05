var gulp = require('gulp'),
    eslint = require('gulp-eslint');
 
gulp.task('lint', function () {
    // ESLint ignores files with "node_modules" paths. 
    // So, it's best to have gulp ignore the directory as well. 
    // Also, Be sure to return the stream from the task; 
    // Otherwise, the task may end before the stream has finished. 
    return gulp.src(['**/*.js','!node_modules/**', '!gulpfile.js'])
        // eslint() attaches the lint output to the "eslint" property 
        // of the file object so it can be used by other modules. 
        .pipe(eslint({
    "ecmaFeatures": {
        "jsx": true,
        "arrowFunctions": true,
        "blockBindings": true,
        "defaultParams": true,
        "destructuring": true,
        "forOf": true,
        "generators": true,
        "objectLiteralComputedProperties": true,
        "objectLiteralShorthandMethods": true,
        "objectLiteralShorthandProperties": true,
        "restParams": true,
        "spread": true,
        "templateStrings": true,
        "modules": true,
        "classes": true
    },
    "env": {
        "browser": true,
        "jasmine": true,
        "node": true
    },
    "rules": {
        "brace-style": 2,
        "camelcase": 2,
        "comma-dangle": [2, "never"],
        "comma-spacing": [2, {
            "before": false,
            "after": true
        }],
        "comma-style": [2, "last"],
        "complexity": [1, 8],
        "consistent-this": [2, "_this"],
        "curly": 2,
        "default-case": 2,
        "dot-notation": 2,
        "eol-last": 2,
        "eqeqeq": 2,
        "guard-for-in": 1,
        "indent": [2, 2, {
            "SwitchCase": 1
        }],
        "key-spacing": [2, {
            "beforeColon": false,
            "afterColon": true
        }],
        "new-cap": 2,
        "new-parens": 2,
        "no-caller": 2,
        "no-debugger": 1,
        "no-dupe-args": 2,
        "no-dupe-keys": 2,
        "no-duplicate-case": 2,
        "no-eq-null": 0,
        "no-eval": 2,
        "no-implied-eval": 2,
        "no-invalid-regexp": 2,
        "no-mixed-spaces-and-tabs": 2,
        "no-redeclare": 2,
        "quote-props": [2, "consistent-as-needed"],
        "no-self-compare": 1,
        "no-shadow-restricted-names": 2,
        "no-trailing-spaces": 2,
        "no-undef": 2,
        "no-undef-init": 2,
        "no-underscore-dangle": 0,
        "no-unreachable": 2,
        "no-unused-vars": 1,
        "no-use-before-define": 2,
        "no-with": 2,
        "one-var": [2, "never"],
        "operator-assignment": [2, "always"],
        "quotes": [2, "single"],
        "radix": 2,
        "semi": [2, "always"],
        "semi-spacing": [2, {
            "before": false,
            "after": true
        }],
        "sort-vars": [1, {
            "ignoreCase": true
        }],
        "space-after-keywords": [2, "always"],
        "space-before-function-paren": [2, {
            "anonymous": "always",
            "named": "never"
        }],
        "space-in-parens": [2, "never"],
        "space-infix-ops": 2,
        "space-unary-ops": [2, {
            "words": true,
            "nonwords": false
        }],
        "strict": [2, "global"],
        "use-isnan": 2,
        "valid-jsdoc": 1,
        "yoda": [2, "never", {
            "exceptRange": false
        }]
    }
}))
        // eslint.format() outputs the lint results to the console. 
        // Alternatively use eslint.formatEach() (see Docs). 
        .pipe(eslint.format())
        // To have the process exit with an error code (1) on 
        // lint error, return the stream and pipe to failAfterError last. 
        .pipe(eslint.failAfterError());
});
 
gulp.task('validate', ['lint'], function () {
    // This will only run if the lint task is successful... 
});