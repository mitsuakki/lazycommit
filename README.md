# `lazycommit`

<p align="center">
  <img alt="demo" src="">
</p>

## About

Tired of the repetitive time-wasting task of thinking up a coherent, short commit message in a readable format that everyone can understand, I decided to leave this thankless task to artificial intelligence.

> The `prepare-commit-msg` hook is run before the commit message editor is fired up but after the default message is created. It lets you edit the default message before the commit author sees it. 
> This hook takes a few parameters: 
> * the path to the file that holds the commit message so far;
> * the type of commit;
> * the commit SHA-1 if this is an amended commit.
>
> This hook generally isn’t useful for normal commits; rather, it’s good for commits where the default message is auto-generated, such as templated commit messages, merge commits, squashed commits, and amended commits.
> You may use it in conjunction with a commit template to programmatically insert information.

## Installation

The pre-compiled binaries can be downloaded from [release page](https://github.com/mitsuakki/lazycommit/releases). <br>
Change the binary permissions to `755` and copy the binary to the sys bin directory.

## Setup

Please first create your OpenAI API key. <br>
The [OpenAPI Platform](https://platform.openai.com/api-keys) allows you to generate a new API key.

## Usage

