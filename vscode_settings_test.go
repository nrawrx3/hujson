package hujson

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var vscodeConfig = `
{
  "editor.fontFamily": "Comic Mono, Gopher Mono, New Heterodox Mono A, Tchig Mono, Iosevka SS16, Maple Mono, Jetbrains Mono NL, Victor Mono, Andale Mono, Fantasque Sans Mono, Menlo, Sometype Mono",
  "editor.tabSize": 8,
  "editor.detectIndentation": false,
  "editor.trimAutoWhitespace": true,
  "editor.formatOnPaste": true,
  "editor.fontWeight": "500",
  "editor.multiCursorModifier": "ctrlCmd",
  "editor.snippetSuggestions": "top",
  "workbench.editor.showTabs": false,
  "workbench.editor.enablePreview": false,
  "terminal.integrated.inheritEnv": false,
  "workbench.startupEditor": "none",
  "[javascript]": {
    "editor.defaultFormatter": "vscode.typescript-language-features",
    "editor.tabSize": 2,
    "editor.useTabStops": false
  },
  "[json]": {
    "editor.defaultFormatter": "vscode.json-language-features"
  },
  "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.suggest.snippetsPreventQuickSuggestions": false,
    "editor.tabSize": 2
  },
  "workbench.editorAssociations": {
    "*.mobileconfig": "default"
  },
  "[python]": {
    "editor.tabSize": 4,
    "editor.defaultFormatter": "ms-python.python"
  },
  "[typescriptreact]": {
    "editor.insertSpaces": true,
    "editor.tabSize": 2,
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[jupyter]": {
    "editor.tabSize": 4
  },
  "editor.fontSize": 14,
  "editor.cursorBlinking": "solid",
  "terminal.integrated.fontFamily": "Gopher Mono",
  "jupyter.askForKernelRestart": false,
  "terminal.integrated.fontSize": 14,
  "go.toolsManagement.autoUpdate": true,
  "[cpp]": {
    "editor.defaultFormatter": "ms-vscode.cpptools"
  },
  "editor.wordWrapColumn": 80,
  "editor.wordWrap": "on",
  "[markdown]": {
    "editor.tabSize": 4,
    "editor.useTabStops": false,
    "editor.defaultFormatter": "yzhang.markdown-all-in-one"
  },
  "[proto3]": {
    "editor.tabSize": 2,
    "editor.useTabStops": false
  },
  "python.formatting.provider": "black",
  "notebook.output.textLineLimit": 1000,
  "insertDateString.formatDate": "YYYY-MM-DDThh:mm:ss+00:00",
  "editor.tokenColorCustomizations": {
    "[Dainty – Nord (chroma 0, lightness 0)]": {
      "comments": "#AAAAF0"
    },
    "[Dainty – Nord (chroma 0, lightness 4)]": {
      "comments": "#AAAAF0"
    },
    "[Dainty – Nord (chroma 4, lightness 3)]": {
      "comments": "#AAAAF0"
    },
    "[Nord]": {
      "comments": "#AAAAF0"
    },
    "[Tame Light (rainglow)]": {
      "comments": "#AAAAF0"
    },
    "[Solarized Yuki]": {
      "comments": "#AAAAF0"
    },
    "[Solarized Neue Bright]": {
      "comments": "#AAAAF0"
    },
    "[Nord Deep]": {
      "comments": "#AAAAF0"
    },
    "[paddy-mist]": {
      "comments": "#AAAAF0"
    },
    "[paddy-mist-upright]": {
      "comments": "#AAAAF0"
    },
    "[Solarized Light]": {
      "comments": "#aa97e2"
    },
    "textMateRules": [
      {
        "scope": [
          "comment",
          "comment.block",
          "comment.block.documentation",
          "comment.line",
          "constant",
          "constant.character",
          "constant.character.escape",
          "constant.numeric",
          "constant.numeric.integer",
          "constant.numeric.float",
          "constant.numeric.hex",
          "constant.numeric.octal",
          "constant.other",
          "constant.regexp",
          "constant.rgb-value",
          "emphasis",
          "entity",
          "entity.name",
          "entity.name.class",
          "entity.name.function",
          "entity.name.method",
          "entity.name.section",
          "entity.name.selector",
          "entity.name.tag",
          "entity.name.type",
          "entity.other",
          "entity.other.attribute-name",
          "entity.other.inherited-class",
          "invalid",
          "invalid.deprecated",
          "invalid.illegal",
          "keyword",
          "keyword.control",
          "keyword.operator",
          "keyword.operator.new",
          "keyword.operator.assignment",
          "keyword.operator.arithmetic",
          "keyword.operator.logical",
          "keyword.other",
          "markup",
          "markup.bold",
          "markup.changed",
          "markup.deleted",
          "markup.heading",
          "markup.inline.raw",
          "markup.inserted",
          "markup.italic",
          "markup.list",
          "markup.list.numbered",
          "markup.list.unnumbered",
          "markup.other",
          "markup.quote",
          "markup.raw",
          "markup.underline",
          "markup.underline.link",
          "meta",
          "meta.block",
          "meta.cast",
          "meta.class",
          "meta.function",
          "meta.function-call",
          "meta.preprocessor",
          "meta.return-type",
          "meta.selector",
          "meta.tag",
          "meta.type.annotation",
          "meta.type",
          "punctuation.definition.string.begin",
          "punctuation.definition.string.end",
          "punctuation.separator",
          "punctuation.separator.continuation",
          "punctuation.terminator",
          "storage",
          "storage.modifier",
          "storage.type",
          "string",
          "string.interpolated",
          "string.other",
          "string.quoted",
          "string.quoted.double",
          "string.quoted.other",
          "string.quoted.single",
          "string.quoted.triple",
          "string.regexp",
          "string.unquoted",
          "strong",
          "support",
          "support.class",
          "support.constant",
          "support.function",
          "support.other",
          "support.type",
          "support.type.property-name",
          "support.variable",
          "variable",
          "variable.language",
          "variable.name",
          "variable.other",
          "variable.other.readwrite",
          "variable.parameter"
        ],
        "settings": {
          "fontStyle": ""
        }
      }
    ]
  },
  "window.nativeTabs": true,
  "editor.minimap.showSlider": "always",
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "better-comments.tags": [
    {
      "tag": "!",
      "color": "#FF2D00",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "?",
      "color": "#3498DB",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "//",
      "color": "#474747",
      "strikethrough": true,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    },
    {
      "tag": "TODO",
      "color": "#be99ff",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "rgba(0.7, 0.7, 0.8, 1.0)",
      "bold": false,
      "italic": false
    },
    {
      "tag": "EXPLAIN",
      "color": "Plum",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "rgba(0.5, 0.5, 0.1, 1.0)",
      "bold": false,
      "italic": false
    },
    {
      "tag": "FILTHY",
      "color": "Red",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "rgba(0.5, 0.5, 0.1, 1.0)",
      "bold": false,
      "italic": false
    },
    {
      "tag": "CONSIDER",
      "color": "Khaki",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "rgba(0.5, 0.5, 0.1, 1.0)",
      "bold": false,
      "italic": false
    },
    {
      "tag": "*",
      "color": "#98C379",
      "strikethrough": false,
      "underline": false,
      "backgroundColor": "transparent",
      "bold": false,
      "italic": false
    }
  ],
  "workbench.iconTheme": "easy-icons",
//   "editor.scrollbar.verticalScrollbarSize": 5,
  "redhat.telemetry.enabled": false,
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.tabSize": 2,
    "editor.insertSpaces": true,
    "editor.formatOnSave": true
  },
  "[yaml]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "files.associations": {
    "*.hujson": "jsonc"
  },
  "json.schemas": [
    {
      "fileMatch": ["*.hujson"],
      "schema": {
        "allowTrailingCommas": true
      }
    }
  ],
"workbench.activityBar.visible": false,
  "editor.inlineSuggest.enabled": true,
  "github.copilot.enable": {
    "*": true,
    "yaml": false,
    "plaintext": false,
    "markdown": false,
    "scminput": false
  },
"editor.minimap.enabled": false,
"git.openRepositoryInParentFolders": "never",
"[dart]": {
        "editor.formatOnSave": true,
        "editor.formatOnType": true,
        "editor.selectionHighlight": false,
        "editor.suggest.snippetsPreventQuickSuggestions": false,
        "editor.suggestSelection": "first",
        "editor.tabCompletion": "onlySnippets",
        "editor.wordBasedSuggestions": false
},
// "editor.cursorSmoothCaretAnimation": "on",
"editor.cursorStyle": "line",
"workbench.colorTheme": "Everforest Dark",
}
`

func TestVscodeSettingsJson(t *testing.T) {
	val, err := Parse([]byte(vscodeConfig))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	packed := val.Pack()

	if diff := cmp.Diff([]byte(vscodeConfig), packed); diff != "" {
		t.Errorf("Packed val mismatch: %s", diff)
	}
}
