package main
type T struct {
	x, y int
	u float32
	_ float32
	A *[]int
	F func()
}

/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiElement(WS_NEW_LINES)('\n')
  TypeDeclarationsImpl
    PsiElement(KEYWORD_TYPE)('type')
    PsiWhiteSpace(' ')
    TypeSpecImpl
      TypeNameDeclaration(T)
        PsiElement(IDENTIFIER)('T')
      PsiWhiteSpace(' ')
      TypeStructImpl
        PsiElement(KEYWORD_STRUCT)('struct')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('x')
          PsiElement(,)(',')
          PsiWhiteSpace(' ')
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('y')
          PsiWhiteSpace(' ')
          TypeNameImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('int')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('u')
          PsiWhiteSpace(' ')
          TypeNameImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('float32')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('_')
          PsiWhiteSpace(' ')
          TypeNameImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('float32')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('A')
          PsiWhiteSpace(' ')
          TypePointerImpl
            PsiElement(*)('*')
            TypeSliceImpl
              PsiElement([)('[')
              PsiElement(])(']')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('int')
        PsiElement(WS_NEW_LINES)('\n')
        PsiWhiteSpace('\t')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('F')
          PsiWhiteSpace(' ')
          TypeFunctionImpl
            PsiElement(KEYWORD_FUNC)('func')
            PsiElement(()('(')
            PsiElement())(')')
        PsiElement(WS_NEW_LINES)('\n')
        PsiElement(})('}')
  PsiElement(WS_NEW_LINES)('\n')