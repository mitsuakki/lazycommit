package llm

var Prompt = `
Hello there ! I have something for you.
You are an expert programmer trying to write a commit message, for that you need to write a summary of the code change.
So, you went over every file that you changed and wrote a summary for each file.

You have to determine the best label for the commit, for that you need to follow some conventions.
Here are the labels you can choose from:

	- feat: A commit of the type feat introduces a new feature to the codebase (this correlates with `MINOR` in Semantic Versioning).
	- fix: A commit of the type fix patches a bug in your codebase (this correlates with `PATCH` in Semantic Versioning).
	- BREAKING CHANGE: A commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
	- docs: Documentation only changes
	- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
	- refactor: A code change that neither fixes a bug nor adds a feature
	- perf: A code change that improves performance
	- test: Adding missing tests or correcting existing tests
	- build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
	- ci: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
	- chore: Other changes that don't modify src or test files
	- revert: Reverts a previous commit

You can also add a scope to your label, for example:
	- feat(parser): add ability to parse arrays

You can also add a body to your commit message, for example:
	- feat(parser): add ability to parse arrays

		The body should include the motivation for the change and contrast this with previous behavior.

You can also add a footer to your commit message, for example:
	- feat(parser): add ability to parse arrays

		BREAKING CHANGE: this change modifies the behavior of the parser to allow parsing arrays.

You can also add issues references to your commit message, for example:
	- feat(parser): add ability to parse arrays

		BREAKING CHANGE: this change modifies the behavior of the parser to allow parsing arrays.

		Closes #123, #245, #992

The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” are to be interpreted as described in RFC 2119.

	1. Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc., followed by the OPTIONAL scope, OPTIONAL !, and REQUIRED terminal colon and space.
	2. The type feat MUST be used when a commit adds a new feature to your application or library.
	3. The type fix MUST be used when a commit represents a bug fix for your application.
	4. A scope MAY be provided after a type. A scope MUST consist of a noun describing a section of the codebase surrounded by parenthesis, e.g., fix(parser):
	5. A description MUST immediately follow the colon and space after the type/scope prefix. The description is a short summary of the code changes, e.g., fix: array parsing issue when multiple spaces were contained in string.
	6. A longer commit body MAY be provided after the short description, providing additional contextual information about the code changes. The body MUST begin one blank line after the description.
	7. A commit body is free-form and MAY consist of any number of newline separated paragraphs.
	8. One or more footers MAY be provided one blank line after the body. Each footer MUST consist of a word token, followed by either a :<space> or <space># separator, followed by a string value (this is inspired by the git trailer convention).
	9. A footer’s token MUST use - in place of whitespace characters, e.g., Acked-by (this helps differentiate the footer section from a multi-paragraph body). An exception is made for BREAKING CHANGE, which MAY also be used as a token.
	10. A footer’s value MAY contain spaces and newlines, and parsing MUST terminate when the next valid footer token/separator pair is observed.
	11. Breaking changes MUST be indicated in the type/scope prefix of a commit, or as an entry in the footer.
	12. If included as a footer, a breaking change MUST consist of the uppercase text BREAKING CHANGE, followed by a colon, space, and description, e.g., BREAKING CHANGE: environment variables now take precedence over config files.
	13. If included in the type/scope prefix, breaking changes MUST be indicated by a ! immediately before the :. If ! is used, BREAKING CHANGE: MAY be omitted from the footer section, and the commit description SHALL be used to describe the breaking change.
	14. Types other than feat and fix MAY be used in your commit messages, e.g., docs: update ref docs.
	15. The units of information that make up Conventional Commits MUST NOT be treated as case sensitive by implementors, with the exception of BREAKING CHANGE which MUST be uppercase.
	16. BREAKING-CHANGE MUST be synonymous with BREAKING CHANGE, when used as a token in a footer.

Now that you know the conventions for label your commit messages, you need some informations before summarizing the code changes.
For every file, there are a few metadata lines, like (for example):
\`\`\`
diff --git a/lib/index.js b/lib/index.js
index aadf691..bfef603 100644
--- a/lib/index.js
+++ b/lib/index.js
\`\`\`

This means that `lib/index.js` was modified in this commit. Note that this is only an example.
Then there is a specifier of the lines that were modified.
A line starting with `+` means it was added.
A line that starting with `-` means that line was deleted.
A line that starts with neither `+` nor `-` is code given for context and better understanding.
It is not part of the diff.
After the git diff of the first file, there will be an empty line, and then the git diff of the next file.

Do not include the file name as another part of the comment.
Do not use the characters `[` or `]` in the summary.
Write every summary comment in a new line.
Comments should be in a bullet point list, each line starting with a `-`.
The summary should not include comments copied from the code.
The output should be easily readable. When in doubt, write less comments and not more. Do not output comments that simply repeat the contents of the file.
Readability is top priority. Write only the most important comments about the diff.

EXAMPLE SUMMARY COMMENTS:

- Raise the amount of returned recordings from `10` to `100`
- Fix a typo in the github action name
- Move the `octokit` initialization to a separate file
- Add an OpenAI API for completions
- Lower numeric tolerance for test files
- Add 2 tests for the inclusive string split function

Most commits will have less comments than this examples list.
The last comment does not include the file names,
because there were more than two relevant files in the hypothetical commit.
Do not include parts of the example in your summary.
It is given only as an example of appropriate comments.

THE FILE SUMMARIES:

{{ .file_changes }}

Now, based on the changes described in the file summaries you can start writing the summary comments for the code changes and choose what's the best label for the commit for each files based on their summaries.
Remember to follow the conventions for commit messages. Don't write changes. 
`